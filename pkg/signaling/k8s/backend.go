package k8s

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"riasc.eu/wice/pkg/crypto"
	"riasc.eu/wice/pkg/pb"
	"riasc.eu/wice/pkg/signaling"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "riasc.eu/wice/pkg/signaling/k8s/apis/wice/v1"
	wicev1 "riasc.eu/wice/pkg/signaling/k8s/client/clientset/versioned"
	informers "riasc.eu/wice/pkg/signaling/k8s/client/informers/externalversions"
)

const (
	cleanupInterval = 1 * time.Minute
	cleanupMaxAge   = 10 * time.Minute
)

type Backend struct {
	signaling.SubscriptionsRegistry

	config BackendConfig

	clientSet *wicev1.Clientset
	informer  cache.SharedInformer

	term chan struct{}

	events chan *pb.Event

	logger *zap.Logger
}

func init() {
	signaling.Backends["k8s"] = &signaling.BackendPlugin{
		New:         NewBackend,
		Description: "Exchange candidates via annotation in Kubernetes Node resource",
	}
}

func NewBackend(cfg *signaling.BackendConfig, events chan *pb.Event) (signaling.Backend, error) {
	var config *rest.Config
	var err error

	b := Backend{
		SubscriptionsRegistry: signaling.NewSubscriptionsRegistry(),
		term:                  make(chan struct{}),
		config:                defaultConfig,
		events:                events,
		logger:                zap.L().Named("backend").With(zap.String("backend", cfg.URI.Scheme)),
	}

	if err := b.config.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}

	if b.config.Kubeconfig == "" {
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		// if you want to change the loading rules (which files in which order), you can do so here

		configOverrides := &clientcmd.ConfigOverrides{}
		// if you want to change override values or bind them to flags, there are methods to help you

		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

		if config, err = kubeConfig.ClientConfig(); err != nil {
			return nil, fmt.Errorf("failed to load config: %w", err)
		}
	} else if b.config.Kubeconfig == "incluster" {
		if config, err = rest.InClusterConfig(); err != nil {
			return nil, fmt.Errorf("failed to get incluster configuration: %w", err)
		}
	} else {
		if config, err = clientcmd.BuildConfigFromFlags("", b.config.Kubeconfig); err != nil {
			return nil, fmt.Errorf("failed to get configuration from flags: %w", err)
		}
	}

	// Create the clientset
	if b.clientSet, err = wicev1.NewForConfig(config); err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	// Create the shared informer factory and use the client to connect to Kubernetes
	factory := informers.NewSharedInformerFactoryWithOptions(b.clientSet, 0, informers.WithNamespace(b.config.Namespace))

	// Get the informer for the right resource, in this case a Pod
	b.informer = factory.Wice().V1().SignalingEnvelopes().Informer()

	b.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    b.onSignalingEnvelopeAdd,
		UpdateFunc: b.onSessionDescriptionUpdate,
	})

	go b.informer.Run(b.term)
	b.logger.Debug("Started watching node resources")

	cache.WaitForNamedCacheSync("signalingenvelopes", b.term, b.informer.HasSynced)

	go b.periodicCleanup()
	b.logger.Debug("Started regular cleanup")

	b.events <- &pb.Event{
		Type: pb.Event_BACKEND_READY,
		Event: &pb.Event_BackendReady{
			BackendReady: &pb.BackendReadyEvent{
				Type: pb.BackendReadyEvent_K8S,
			},
		},
	}

	return &b, nil
}

func (b *Backend) Subscribe(kp *crypto.KeyPair) (chan *pb.SignalingMessage, error) {
	b.logger.Info("Subscribe to messages from peer", zap.Any("kp", kp))

	sub, err := b.NewSubscription(kp)
	if err != nil {
		return nil, fmt.Errorf("failed create subscription: %w", err)
	}

	// Process existing envelopes in cache
	if err := b.processByKeyPair(kp); err != nil {
		return nil, err
	}

	return sub.Channel, nil
}

func (b *Backend) Publish(kp *crypto.KeyPair, msg *pb.SignalingMessage) error {
	var err error

	b.logger.Debug("Published signaling message",
		zap.Any("kp", kp),
		zap.Any("msg", msg),
	)

	envs := b.clientSet.WiceV1().SignalingEnvelopes(b.config.Namespace)

	pbEnv, err := msg.Encrypt(kp)
	if err != nil {
		return fmt.Errorf("failed to encrypt message: %w", err)
	}

	env := &v1.SignalingEnvelope{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: b.config.GenerateName,
		},
	}

	pbEnv.DeepCopyInto(&env.SignalingEnvelope)

	if env, err = envs.Create(context.Background(), env, metav1.CreateOptions{}); err != nil {
		return fmt.Errorf("failed to create envelope: %w", err)
	}

	b.logger.Debug("Created new SignalingEnvelope", zap.String("name", env.ObjectMeta.Name))

	return nil
}

func (b *Backend) Close() error {
	close(b.term)

	return nil // TODO
}

func (b *Backend) onSignalingEnvelopeAdd(obj interface{}) {
	env := obj.(*v1.SignalingEnvelope)

	b.logger.Debug("SignalingEnvelope added", zap.String("name", env.ObjectMeta.Name))
	if err := b.process(env); err != nil {
		b.logger.Error("Failed to process SignalEnvelope", zap.Error(err))
	}
}

func (b *Backend) onSessionDescriptionUpdate(_ interface{}, new interface{}) {
	newEnv := new.(*v1.SignalingEnvelope)

	b.logger.Debug("SignalingEnvelope updated", zap.String("name", newEnv.ObjectMeta.Name))
	if err := b.process(newEnv); err != nil {
		b.logger.Error("Failed to process SignalEnvelope", zap.Error(err))
	}
}

func (b *Backend) process(env *v1.SignalingEnvelope) error {
	sender, err := crypto.ParseKeyBytes(env.Sender)
	if err != nil {
		return fmt.Errorf("invalid key: %w", err)
	}

	sub, err := b.GetSubscription(&sender)
	if err != nil {
		return nil // ignore envelopes not addressed to us
	}

	if err := sub.NewMessage(&env.SignalingEnvelope); err != nil {
		return err
	}

	// Delete envelope
	envs := b.clientSet.WiceV1().SignalingEnvelopes(b.config.Namespace)
	if err := envs.Delete(context.Background(), env.ObjectMeta.Name, metav1.DeleteOptions{}); err != nil {
		b.logger.Warn("Failed to delete envelope", zap.Error(err))
	} else {
		b.logger.Debug("Deleted envelope", zap.String("envelope", env.ObjectMeta.Name))
	}

	return nil
}

func (b *Backend) processByKeyPair(kp *crypto.KeyPair) error {
	store := b.informer.GetStore()
	for _, obj := range store.List() {
		if env, ok := obj.(*v1.SignalingEnvelope); ok {
			if err := b.process(env); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Backend) periodicCleanup() {
	ticker := time.NewTicker(cleanupInterval)

	b.cleanup()
	for range ticker.C {
		b.cleanup()
	}
}

func (b *Backend) cleanup() {
	store := b.informer.GetStore()
	envs := b.clientSet.WiceV1().SignalingEnvelopes(b.config.Namespace)

	for _, obj := range store.List() {
		if env, ok := obj.(*v1.SignalingEnvelope); ok {
			if time.Since(env.ObjectMeta.CreationTimestamp.Time) > cleanupMaxAge {
				if err := envs.Delete(context.Background(), env.ObjectMeta.Name, metav1.DeleteOptions{}); err != nil {
					b.logger.Error("Failed to delete envelope", zap.Any("name", env.ObjectMeta.Name), zap.Error(err))
				} else {
					b.logger.Debug("Deleted stale SignalingEnvelope", zap.String("name", env.ObjectMeta.Name))
				}
			}
		}
	}
}
