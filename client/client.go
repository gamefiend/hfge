package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Client type holds connection data to a HFGE server
type Client struct {
	ServerAddress string
}

type Result struct{
	Hex		int		`json:"current_hex"`
	Content	string	`json:"content"`
}

// New creates a new HFGE Client
func New(serverAddress string) (*Client, error) {
	if serverAddress == "" {
		return nil, fmt.Errorf("bad server address %q (must not be empty)", serverAddress)
	}

	return &Client{
		ServerAddress: serverAddress,
	}, nil
}

// List returns a list of hexflowers from an hfge server
func (c *Client) List() (string, error) {
	listURL := fmt.Sprintf("%s/list", c.ServerAddress)
	resp, err := http.Get(listURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("Server returned status code %d", resp.StatusCode)
	}
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (c *Client) Start(content string) (*Result, error) {
	result := &Result{}
	startURL := fmt.Sprintf("%s/%s/", c.ServerAddress, content)
	resp, err := http.Get(startURL)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK
		return nil, errors.Errorf("Server returned status code %d", resp.StatusCode)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, result)
	if err != nil{
		return nil, err
	}
	return result, nil
}



