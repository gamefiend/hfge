package server_test

import (
	"fmt"
	"hex/server"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetContent(t *testing.T) {
	ns := server.New(":8088")
	ns.Start()
	defer ns.Stop()
	tcs := []string{
		"terrain",
		"weather",
	}
	for _, tc := range tcs {
		response, err := http.Get(fmt.Sprintf("http://localhost:8088/", tc))
		if err != nil {
			t.Fatalf("could not connect: %v", err)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatalf("unexpected response status: %v", response.StatusCode)
		}
		defer response.Body.Close()
		got, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatalf("problems reading response body : %v", err)
		}
		want, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", tc))
		if err != nil {
			t.Fatalf("problems reading test file: %v", err)
		}
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(string(want), string(got)))
		}
	}
}
