package client_test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
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
	want := "terrain\nweather\n"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.URL.Path != "/list" {
			t.Errorf("want URL /list, but got %q", r.URL.Path)
		}
		fmt.Fprint(w, want)
	}))
	c, err := client.New(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.List()
	if err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Fatal("server not called")
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestClientStart(t *testing.T) {
	var called bool
	wantHex := 10
	wantContent := "special"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.URL.Path != "/terrain/" {
			t.Errorf("want URL /terrain/, but got %q", r.URL.Path)
		}
		fmt.Fprint(w, `{
"current_hex": 10,
"content": "special"
}`)
	}))
	c, err := client.New(ts.URL)
	if err != nil{
		t.Fatal(err)
	}
	result, err := c.Start("terrain")
	if !called {
		t.Fatal("server not called")
	}
	if !cmp.Equal(wantHex, result.Hex) {
		t.Error(cmp.Diff(wantHex, result.Hex))
	}
	if !cmp.Equal(wantContent, result.Content){
		t.Error(cmp.Diff(wantContent, result.Content))
	}


}
