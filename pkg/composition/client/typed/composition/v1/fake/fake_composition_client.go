// Generated code
// run `make generate` to update

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/atlassian/voyager/pkg/composition/client/typed/composition/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCompositionV1 struct {
	*testing.Fake
}

func (c *FakeCompositionV1) ServiceDescriptors() v1.ServiceDescriptorInterface {
	return &FakeServiceDescriptors{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCompositionV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
