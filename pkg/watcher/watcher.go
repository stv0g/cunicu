// Package watcher keeps track and monitors for new, removed and modified WireGuard interfaces and peers.
package watcher

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/stv0g/cunicu/pkg/core"
	"github.com/stv0g/cunicu/pkg/crypto"
	"github.com/stv0g/cunicu/pkg/log"
	"github.com/stv0g/cunicu/pkg/util"
	"go.uber.org/zap"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var errNotSupported = errors.New("not supported on this platform")

const (
	InterfaceAdded InterfaceEventOp = iota
	InterfaceDeleted
)

type InterfaceFilterFunc func(string) bool

type (
	InterfaceEventOp int
	InterfaceEvent   struct {
		Op   InterfaceEventOp
		Name string
	}
)

func (ls InterfaceEventOp) String() string {
	switch ls {
	case InterfaceAdded:
		return "added"
	case InterfaceDeleted:
		return "deleted"
	default:
		return ""
	}
}

func (e InterfaceEvent) String() string {
	return fmt.Sprintf("%s %s", e.Name, e.Op)
}

// Watcher monitors both userspace and kernel for changes to WireGuard interfaces
type Watcher struct {
	interfaces core.InterfaceList
	devices    []*wgtypes.Device

	mu sync.RWMutex

	onInterface []core.InterfaceHandler

	client *wgctrl.Client

	events chan InterfaceEvent
	errors chan error
	stop   chan any
	manual chan any

	// Settings
	filter   InterfaceFilterFunc
	interval time.Duration

	logger *zap.Logger
}

func New(client *wgctrl.Client, interval time.Duration, filter InterfaceFilterFunc) (*Watcher, error) {
	return &Watcher{
		interfaces: core.InterfaceList{},
		devices:    []*wgtypes.Device{},

		onInterface: []core.InterfaceHandler{},

		client:   client,
		filter:   filter,
		interval: interval,

		events: make(chan InterfaceEvent, 16),
		errors: make(chan error, 16),
		stop:   make(chan any),
		manual: make(chan any, 16),

		logger: zap.L().Named("watcher"),
	}, nil
}

func (w *Watcher) Close() error {
	if err := w.Sync(); err != nil {
		return fmt.Errorf("final sync failed: %w", err)
	}

	close(w.stop)

	return nil
}

func (w *Watcher) Watch() {
	if err := w.watchUser(); err != nil {
		w.logger.Fatal("Failed to watch userspace interfaces", zap.Error(err))
	}
	w.logger.Debug("Started watching for changes of WireGuard userspace devices")

	if err := w.watchKernel(); err != nil && !errors.Is(err, errNotSupported) {
		w.logger.Fatal("Failed to watch kernel interfaces", zap.Error(err))
	}
	w.logger.Debug("Started watching for changes of WireGuard kernel devices")

	ticker := &time.Ticker{}
	if w.interval > 0 {
		ticker = time.NewTicker(w.interval)
		defer ticker.Stop()
	}

	logger := w.logger.WithOptions(log.WithVerbose(10))

out:
	for {
		select {
		case <-w.manual:
			logger.Debug("Start manual interface synchronization")
			if err := w.sync(); err != nil {
				w.logger.Error("Synchronization failed", zap.Error(err))
			}

		// We still a need periodic sync we can not (yet) monitor WireGuard interfaces
		// for changes via a netlink socket (patch is pending)
		case <-ticker.C:
			logger.Debug("Start periodic interface synchronization")
			if err := w.sync(); err != nil {
				w.logger.Error("Synchronization failed", zap.Error(err))
			}

		case event := <-w.events:
			logger.Debug("Received interface event", zap.Any("event", event))
			if err := w.sync(); err != nil {
				w.logger.Error("Synchronization failed", zap.Error(err))
			}

		case err := <-w.errors:
			w.logger.Error("Failed to watch for interface changes", zap.Error(err))

		case <-w.stop:
			break out
		}
	}
}

func (w *Watcher) Sync() error {
	w.manual <- nil

	return nil
}

func (w *Watcher) sync() error {
	var err error
	var newDevs []*wgtypes.Device
	oldDevs := w.devices

	w.mu.Lock()

	if newDevs, err = w.client.Devices(); err != nil {
		w.mu.Unlock()
		return fmt.Errorf("failed to list WireGuard interfaces: %w", err)
	}

	// Ignore devices which do not match the filter
	newDevs = util.SliceFilter(newDevs, func(d *wgtypes.Device) bool {
		return w.filter == nil || w.filter(d.Name)
	})

	added, removed, kept := util.SliceDiffFunc(oldDevs, newDevs, func(a, b *wgtypes.Device) int {
		return strings.Compare(a.Name, b.Name)
	})

	w.mu.Unlock()

	for _, wgd := range removed {
		i, ok := w.interfaces[wgd.Name]
		if !ok {
			w.logger.Warn("Failed to find matching interface", zap.Any("intf", wgd.Name))
			continue
		}

		w.logger.Info("Interface removed", zap.String("intf", wgd.Name))

		for _, h := range w.onInterface {
			h.OnInterfaceRemoved(i)
		}

		delete(w.interfaces, wgd.Name)
	}

	for _, wgd := range added {
		w.logger.Info("Interface added", zap.String("intf", wgd.Name))

		i, err := core.NewInterface(wgd, w.client)
		if err != nil {
			w.logger.Fatal("Failed to create new interface",
				zap.Error(err),
				zap.String("intf", wgd.Name),
			)
		}

		// We purposefully prune the peer list here to force full initial sync of all peers
		i.Device.Peers = nil

		w.interfaces[wgd.Name] = i

		for _, h := range w.onInterface {
			h.OnInterfaceAdded(i)
		}

		i.Sync(wgd)
	}

	for _, wgd := range kept {
		i, ok := w.interfaces[wgd.Name]
		if !ok {
			w.logger.Warn("Failed to find matching interface", zap.Any("intf", wgd.Name))
			continue
		}

		i.Sync(wgd)
	}

	w.devices = newDevs

	return nil
}

func (w *Watcher) Peer(intf string, pk *crypto.Key) *core.Peer {
	i := w.InterfaceByName(intf)
	if i == nil {
		return nil
	}

	if p, ok := i.Peers[*pk]; ok {
		return p
	}

	return nil
}

func (w *Watcher) PeerByPublicKey(pk *crypto.Key) *core.Peer {
	w.mu.RLock()
	defer w.mu.RUnlock()

	for _, i := range w.interfaces {
		if p, ok := i.Peers[*pk]; ok {
			return p
		}
	}

	return nil
}

func (w *Watcher) InterfaceByName(name string) *core.Interface {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.interfaces.ByName(name)
}

func (w *Watcher) InterfaceByPublicKey(pk crypto.Key) *core.Interface {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.interfaces.ByPublicKey(pk)
}

func (w *Watcher) InterfaceByIndex(idx int) *core.Interface {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.interfaces.ByIndex(idx)
}

func (w *Watcher) ForEachInterface(cb func(i *core.Interface) error) error {
	w.mu.RLock()
	defer w.mu.RUnlock()

	for _, i := range w.interfaces {
		if err := cb(i); err != nil {
			return err
		}
	}

	return nil
}

func (w *Watcher) ForEachPeer(cb func(p *core.Peer) error) error {
	return w.ForEachInterface(func(i *core.Interface) error {
		for _, p := range i.Peers {
			if err := cb(p); err != nil {
				return err
			}
		}

		return nil
	})
}
