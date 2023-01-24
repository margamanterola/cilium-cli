// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumLoadBalancerIPPoolsGetter has a method to return a CiliumLoadBalancerIPPoolInterface.
// A group's client should implement this interface.
type CiliumLoadBalancerIPPoolsGetter interface {
	CiliumLoadBalancerIPPools() CiliumLoadBalancerIPPoolInterface
}

// CiliumLoadBalancerIPPoolInterface has methods to work with CiliumLoadBalancerIPPool resources.
type CiliumLoadBalancerIPPoolInterface interface {
	Create(ctx context.Context, ciliumLoadBalancerIPPool *v2alpha1.CiliumLoadBalancerIPPool, opts v1.CreateOptions) (*v2alpha1.CiliumLoadBalancerIPPool, error)
	Update(ctx context.Context, ciliumLoadBalancerIPPool *v2alpha1.CiliumLoadBalancerIPPool, opts v1.UpdateOptions) (*v2alpha1.CiliumLoadBalancerIPPool, error)
	UpdateStatus(ctx context.Context, ciliumLoadBalancerIPPool *v2alpha1.CiliumLoadBalancerIPPool, opts v1.UpdateOptions) (*v2alpha1.CiliumLoadBalancerIPPool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.CiliumLoadBalancerIPPool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.CiliumLoadBalancerIPPoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumLoadBalancerIPPool, err error)
	CiliumLoadBalancerIPPoolExpansion
}

// ciliumLoadBalancerIPPools implements CiliumLoadBalancerIPPoolInterface
type ciliumLoadBalancerIPPools struct {
	client rest.Interface
}

// newCiliumLoadBalancerIPPools returns a CiliumLoadBalancerIPPools
func newCiliumLoadBalancerIPPools(c *CiliumV2alpha1Client) *ciliumLoadBalancerIPPools {
	return &ciliumLoadBalancerIPPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumLoadBalancerIPPool, and returns the corresponding ciliumLoadBalancerIPPool object, and an error if there is any.
func (c *ciliumLoadBalancerIPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumLoadBalancerIPPool{}
	err = c.client.Get().
		Resource("ciliumloadbalancerippools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumLoadBalancerIPPools that match those selectors.
func (c *ciliumLoadBalancerIPPools) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumLoadBalancerIPPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.CiliumLoadBalancerIPPoolList{}
	err = c.client.Get().
		Resource("ciliumloadbalancerippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumLoadBalancerIPPools.
func (c *ciliumLoadBalancerIPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumloadbalancerippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumLoadBalancerIPPool and creates it.  Returns the server's representation of the ciliumLoadBalancerIPPool, and an error, if there is any.
func (c *ciliumLoadBalancerIPPools) Create(ctx context.Context, ciliumLoadBalancerIPPool *v2alpha1.CiliumLoadBalancerIPPool, opts v1.CreateOptions) (result *v2alpha1.CiliumLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumLoadBalancerIPPool{}
	err = c.client.Post().
		Resource("ciliumloadbalancerippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumLoadBalancerIPPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumLoadBalancerIPPool and updates it. Returns the server's representation of the ciliumLoadBalancerIPPool, and an error, if there is any.
func (c *ciliumLoadBalancerIPPools) Update(ctx context.Context, ciliumLoadBalancerIPPool *v2alpha1.CiliumLoadBalancerIPPool, opts v1.UpdateOptions) (result *v2alpha1.CiliumLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumLoadBalancerIPPool{}
	err = c.client.Put().
		Resource("ciliumloadbalancerippools").
		Name(ciliumLoadBalancerIPPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumLoadBalancerIPPool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ciliumLoadBalancerIPPools) UpdateStatus(ctx context.Context, ciliumLoadBalancerIPPool *v2alpha1.CiliumLoadBalancerIPPool, opts v1.UpdateOptions) (result *v2alpha1.CiliumLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumLoadBalancerIPPool{}
	err = c.client.Put().
		Resource("ciliumloadbalancerippools").
		Name(ciliumLoadBalancerIPPool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumLoadBalancerIPPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumLoadBalancerIPPool and deletes it. Returns an error if one occurs.
func (c *ciliumLoadBalancerIPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumloadbalancerippools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumLoadBalancerIPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumloadbalancerippools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumLoadBalancerIPPool.
func (c *ciliumLoadBalancerIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumLoadBalancerIPPool{}
	err = c.client.Patch(pt).
		Resource("ciliumloadbalancerippools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
