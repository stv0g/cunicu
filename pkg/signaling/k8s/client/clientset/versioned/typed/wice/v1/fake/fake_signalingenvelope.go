/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	wicev1 "riasc.eu/wice/pkg/signaling/k8s/apis/wice/v1"
)

// FakeSignalingEnvelopes implements SignalingEnvelopeInterface
type FakeSignalingEnvelopes struct {
	Fake *FakeWiceV1
	ns   string
}

var signalingenvelopesResource = schema.GroupVersionResource{Group: "wice.riasc.eu", Version: "v1", Resource: "signalingenvelopes"}

var signalingenvelopesKind = schema.GroupVersionKind{Group: "wice.riasc.eu", Version: "v1", Kind: "SignalingEnvelope"}

// Get takes name of the signalingEnvelope, and returns the corresponding signalingEnvelope object, and an error if there is any.
func (c *FakeSignalingEnvelopes) Get(ctx context.Context, name string, options v1.GetOptions) (result *wicev1.SignalingEnvelope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(signalingenvelopesResource, c.ns, name), &wicev1.SignalingEnvelope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wicev1.SignalingEnvelope), err
}

// List takes label and field selectors, and returns the list of SignalingEnvelopes that match those selectors.
func (c *FakeSignalingEnvelopes) List(ctx context.Context, opts v1.ListOptions) (result *wicev1.SignalingEnvelopeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(signalingenvelopesResource, signalingenvelopesKind, c.ns, opts), &wicev1.SignalingEnvelopeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &wicev1.SignalingEnvelopeList{ListMeta: obj.(*wicev1.SignalingEnvelopeList).ListMeta}
	for _, item := range obj.(*wicev1.SignalingEnvelopeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested signalingEnvelopes.
func (c *FakeSignalingEnvelopes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(signalingenvelopesResource, c.ns, opts))

}

// Create takes the representation of a signalingEnvelope and creates it.  Returns the server's representation of the signalingEnvelope, and an error, if there is any.
func (c *FakeSignalingEnvelopes) Create(ctx context.Context, signalingEnvelope *wicev1.SignalingEnvelope, opts v1.CreateOptions) (result *wicev1.SignalingEnvelope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(signalingenvelopesResource, c.ns, signalingEnvelope), &wicev1.SignalingEnvelope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wicev1.SignalingEnvelope), err
}

// Update takes the representation of a signalingEnvelope and updates it. Returns the server's representation of the signalingEnvelope, and an error, if there is any.
func (c *FakeSignalingEnvelopes) Update(ctx context.Context, signalingEnvelope *wicev1.SignalingEnvelope, opts v1.UpdateOptions) (result *wicev1.SignalingEnvelope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(signalingenvelopesResource, c.ns, signalingEnvelope), &wicev1.SignalingEnvelope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wicev1.SignalingEnvelope), err
}

// Delete takes name of the signalingEnvelope and deletes it. Returns an error if one occurs.
func (c *FakeSignalingEnvelopes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(signalingenvelopesResource, c.ns, name, opts), &wicev1.SignalingEnvelope{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSignalingEnvelopes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(signalingenvelopesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &wicev1.SignalingEnvelopeList{})
	return err
}

// Patch applies the patch and returns the patched signalingEnvelope.
func (c *FakeSignalingEnvelopes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *wicev1.SignalingEnvelope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(signalingenvelopesResource, c.ns, name, pt, data, subresources...), &wicev1.SignalingEnvelope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*wicev1.SignalingEnvelope), err
}
