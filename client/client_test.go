package client_test

import (
	"bytes"
	"hex/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewReturnsErrorForInvalidServer(t *testing.T) {
	_, err := client.New("")
	if err == nil {
		t.Error("want error for invalid server address, but got nil")
	}
}

func TestNewReturnsClientForValidServer(t *testing.T) {
	var c *client.Client
	c, err := client.New("literally anything that's not empty string")
	if err != nil {
		t.Errorf("want no error for valid server address, but got %v", err)
	}
	if c == nil {
		t.Error("want non-nil Client pointer for valid server address, got nil")
	}
}

func TestListCallsListEndpoint(t *testing.T) {
	var called bool
	ts := httptest.NewTLSServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		called = true
		if r.URL.RawPath != "/list" {
			t.Errorf("want URL /list, but got %q", r.URL.RawPath)
		}
	}))
	c, err := client.New(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Fatal("server not called")
	}
	output := bytes.Buffer{}
	err = c.List(&output)
	if err != nil {
		t.Fatal(err)
	}
}