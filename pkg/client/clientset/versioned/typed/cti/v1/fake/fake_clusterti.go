/*
Copyright 2018 The Kubernetes Authors.

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
	ctiv1 "github.com/IBM/trusted-service-identity/pkg/apis/cti/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTIs implements ClusterTIInterface
type FakeClusterTIs struct {
	Fake *FakeTrustedV1
	ns   string
}

var clustertisResource = schema.GroupVersionResource{Group: "trusted.identity", Version: "v1", Resource: "clustertis"}

var clustertisKind = schema.GroupVersionKind{Group: "trusted.identity", Version: "v1", Kind: "ClusterTI"}

// Get takes name of the clusterTI, and returns the corresponding clusterTI object, and an error if there is any.
func (c *FakeClusterTIs) Get(name string, options v1.GetOptions) (result *ctiv1.ClusterTI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertisResource, c.ns, name), &ctiv1.ClusterTI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ctiv1.ClusterTI), err
}

// List takes label and field selectors, and returns the list of ClusterTIs that match those selectors.
func (c *FakeClusterTIs) List(opts v1.ListOptions) (result *ctiv1.ClusterTIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertisResource, clustertisKind, c.ns, opts), &ctiv1.ClusterTIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ctiv1.ClusterTIList{}
	for _, item := range obj.(*ctiv1.ClusterTIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTIs.
func (c *FakeClusterTIs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustertisResource, c.ns, opts))

}

// Create takes the representation of a clusterTI and creates it.  Returns the server's representation of the clusterTI, and an error, if there is any.
func (c *FakeClusterTIs) Create(clusterTI *ctiv1.ClusterTI) (result *ctiv1.ClusterTI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertisResource, c.ns, clusterTI), &ctiv1.ClusterTI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ctiv1.ClusterTI), err
}

// Update takes the representation of a clusterTI and updates it. Returns the server's representation of the clusterTI, and an error, if there is any.
func (c *FakeClusterTIs) Update(clusterTI *ctiv1.ClusterTI) (result *ctiv1.ClusterTI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertisResource, c.ns, clusterTI), &ctiv1.ClusterTI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ctiv1.ClusterTI), err
}

// Delete takes name of the clusterTI and deletes it. Returns an error if one occurs.
func (c *FakeClusterTIs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustertisResource, c.ns, name), &ctiv1.ClusterTI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTIs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertisResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ctiv1.ClusterTIList{})
	return err
}

// Patch applies the patch and returns the patched clusterTI.
func (c *FakeClusterTIs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ctiv1.ClusterTI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertisResource, c.ns, name, data, subresources...), &ctiv1.ClusterTI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ctiv1.ClusterTI), err
}
