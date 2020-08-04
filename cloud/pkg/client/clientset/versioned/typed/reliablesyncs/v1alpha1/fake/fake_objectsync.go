/*
Copyright 2020 The KubeEdge Authors.

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

	v1alpha1 "github.com/kubeedge/kubeedge/cloud/pkg/apis/reliablesyncs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeObjectSyncs implements ObjectSyncInterface
type FakeObjectSyncs struct {
	Fake *FakeReliablesyncsV1alpha1
	ns   string
}

var objectsyncsResource = schema.GroupVersionResource{Group: "reliablesyncs.kubeedge.io", Version: "v1alpha1", Resource: "objectsyncs"}

var objectsyncsKind = schema.GroupVersionKind{Group: "reliablesyncs.kubeedge.io", Version: "v1alpha1", Kind: "ObjectSync"}

// Get takes name of the objectSync, and returns the corresponding objectSync object, and an error if there is any.
func (c *FakeObjectSyncs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ObjectSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objectsyncsResource, c.ns, name), &v1alpha1.ObjectSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectSync), err
}

// List takes label and field selectors, and returns the list of ObjectSyncs that match those selectors.
func (c *FakeObjectSyncs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ObjectSyncList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objectsyncsResource, objectsyncsKind, c.ns, opts), &v1alpha1.ObjectSyncList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ObjectSyncList{ListMeta: obj.(*v1alpha1.ObjectSyncList).ListMeta}
	for _, item := range obj.(*v1alpha1.ObjectSyncList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectSyncs.
func (c *FakeObjectSyncs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objectsyncsResource, c.ns, opts))

}

// Create takes the representation of a objectSync and creates it.  Returns the server's representation of the objectSync, and an error, if there is any.
func (c *FakeObjectSyncs) Create(ctx context.Context, objectSync *v1alpha1.ObjectSync, opts v1.CreateOptions) (result *v1alpha1.ObjectSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objectsyncsResource, c.ns, objectSync), &v1alpha1.ObjectSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectSync), err
}

// Update takes the representation of a objectSync and updates it. Returns the server's representation of the objectSync, and an error, if there is any.
func (c *FakeObjectSyncs) Update(ctx context.Context, objectSync *v1alpha1.ObjectSync, opts v1.UpdateOptions) (result *v1alpha1.ObjectSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objectsyncsResource, c.ns, objectSync), &v1alpha1.ObjectSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectSync), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectSyncs) UpdateStatus(ctx context.Context, objectSync *v1alpha1.ObjectSync, opts v1.UpdateOptions) (*v1alpha1.ObjectSync, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(objectsyncsResource, "status", c.ns, objectSync), &v1alpha1.ObjectSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectSync), err
}

// Delete takes name of the objectSync and deletes it. Returns an error if one occurs.
func (c *FakeObjectSyncs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objectsyncsResource, c.ns, name), &v1alpha1.ObjectSync{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectSyncs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objectsyncsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ObjectSyncList{})
	return err
}

// Patch applies the patch and returns the patched objectSync.
func (c *FakeObjectSyncs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectSync, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objectsyncsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ObjectSync{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectSync), err
}
