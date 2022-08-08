package osc

import goosc "github.com/hypebeast/go-osc/osc"

type Client struct {
	client *goosc.Client
}

func NewClient(client *goosc.Client) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Client() *goosc.Client {
	return c.client
}
