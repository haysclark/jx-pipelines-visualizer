// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReleasesGetter has a method to return a ReleaseInterface.
// A group's client should implement this interface.
type ReleasesGetter interface {
	Releases(namespace string) ReleaseInterface
}

// ReleaseInterface has methods to work with Release resources.
type ReleaseInterface interface {
	Create(*v1.Release) (*v1.Release, error)
	Update(*v1.Release) (*v1.Release, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Release, error)
	List(opts metav1.ListOptions) (*v1.ReleaseList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Release, err error)
	ReleaseExpansion
}

// releases implements ReleaseInterface
type releases struct {
	client rest.Interface
	ns     string
}

// newReleases returns a Releases
func newReleases(c *JenkinsV1Client, namespace string) *releases {
	return &releases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the release, and returns the corresponding release object, and an error if there is any.
func (c *releases) Get(name string, options metav1.GetOptions) (result *v1.Release, err error) {
	result = &v1.Release{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("releases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Releases that match those selectors.
func (c *releases) List(opts metav1.ListOptions) (result *v1.ReleaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ReleaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("releases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested releases.
func (c *releases) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("releases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a release and creates it.  Returns the server's representation of the release, and an error, if there is any.
func (c *releases) Create(release *v1.Release) (result *v1.Release, err error) {
	result = &v1.Release{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("releases").
		Body(release).
		Do().
		Into(result)
	return
}

// Update takes the representation of a release and updates it. Returns the server's representation of the release, and an error, if there is any.
func (c *releases) Update(release *v1.Release) (result *v1.Release, err error) {
	result = &v1.Release{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("releases").
		Name(release.Name).
		Body(release).
		Do().
		Into(result)
	return
}

// Delete takes name of the release and deletes it. Returns an error if one occurs.
func (c *releases) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("releases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *releases) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("releases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched release.
func (c *releases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Release, err error) {
	result = &v1.Release{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("releases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
