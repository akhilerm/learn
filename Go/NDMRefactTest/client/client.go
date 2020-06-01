package client

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	cfg *rest.Config

	// client is the controller-runtime client used to interface with etcd
	client client.Client

	// namespace in which this client is operating
	namespace string
}

func New() (*Client, error) {
	// methods to initialize the client
	cl := &Client{}
	return cl, nil
}

func (c *Client) setNamespace() {

}

func (c *Client) SetNamespace() {

}