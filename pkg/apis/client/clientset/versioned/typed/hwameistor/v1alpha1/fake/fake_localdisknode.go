// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLocalDiskNodes implements LocalDiskNodeInterface
type FakeLocalDiskNodes struct {
	Fake *FakeHwameistorV1alpha1
}

var localdisknodesResource = schema.GroupVersionResource{Group: "hwameistor.io", Version: "v1alpha1", Resource: "localdisknodes"}

var localdisknodesKind = schema.GroupVersionKind{Group: "hwameistor.io", Version: "v1alpha1", Kind: "LocalDiskNode"}

// Get takes name of the localDiskNode, and returns the corresponding localDiskNode object, and an error if there is any.
func (c *FakeLocalDiskNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalDiskNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localdisknodesResource, name), &v1alpha1.LocalDiskNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalDiskNode), err
}

// List takes label and field selectors, and returns the list of LocalDiskNodes that match those selectors.
func (c *FakeLocalDiskNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalDiskNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localdisknodesResource, localdisknodesKind, opts), &v1alpha1.LocalDiskNodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocalDiskNodeList{ListMeta: obj.(*v1alpha1.LocalDiskNodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocalDiskNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localDiskNodes.
func (c *FakeLocalDiskNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localdisknodesResource, opts))
}

// Create takes the representation of a localDiskNode and creates it.  Returns the server's representation of the localDiskNode, and an error, if there is any.
func (c *FakeLocalDiskNodes) Create(ctx context.Context, localDiskNode *v1alpha1.LocalDiskNode, opts v1.CreateOptions) (result *v1alpha1.LocalDiskNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localdisknodesResource, localDiskNode), &v1alpha1.LocalDiskNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalDiskNode), err
}

// Update takes the representation of a localDiskNode and updates it. Returns the server's representation of the localDiskNode, and an error, if there is any.
func (c *FakeLocalDiskNodes) Update(ctx context.Context, localDiskNode *v1alpha1.LocalDiskNode, opts v1.UpdateOptions) (result *v1alpha1.LocalDiskNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localdisknodesResource, localDiskNode), &v1alpha1.LocalDiskNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalDiskNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalDiskNodes) UpdateStatus(ctx context.Context, localDiskNode *v1alpha1.LocalDiskNode, opts v1.UpdateOptions) (*v1alpha1.LocalDiskNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localdisknodesResource, "status", localDiskNode), &v1alpha1.LocalDiskNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalDiskNode), err
}

// Delete takes name of the localDiskNode and deletes it. Returns an error if one occurs.
func (c *FakeLocalDiskNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(localdisknodesResource, name), &v1alpha1.LocalDiskNode{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalDiskNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localdisknodesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocalDiskNodeList{})
	return err
}

// Patch applies the patch and returns the patched localDiskNode.
func (c *FakeLocalDiskNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalDiskNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localdisknodesResource, name, pt, data, subresources...), &v1alpha1.LocalDiskNode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalDiskNode), err
}