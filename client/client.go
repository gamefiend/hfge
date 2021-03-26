package client

import (
	"fmt"
	"io"
)

type Client struct{}

func New(serverAddress string) (*Client, error) {
	if serverAddress == "" {
		return nil, fmt.Errorf("bad server address %q (must not be empty)", serverAddress)
	}
	return &Client{}, nil
}

func (c *Client) List(w io.Writer) error {
	return nil
}