package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Client type holds connection data to a HFGE server
type Client struct {
	serverAddress string
}

// New creates a new HFGE Client
func New(serverAddress string) (*Client, error) {
	if serverAddress == "" {
		return nil, fmt.Errorf("bad server address %q (must not be empty)", serverAddress)
	}
	resp, err := http.Get(fmt.Sprintf("%s/ok", serverAddress))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	// fmt.Println(resp.Body)
	return &Client{
		serverAddress: serverAddress,
	}, nil
}

// List returns a list of hexflowers from an hfge server
func (c *Client) List(w io.Writer) (string, error) {
	listUrl := fmt.Sprintf("%s/list", c.serverAddress)
	resp, err := http.Get(listUrl)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("Server returned status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func responseBodyToString(response io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String()
}
