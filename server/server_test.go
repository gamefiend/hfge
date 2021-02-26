package server_test

import (
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)
import "hex/server"

func TestGetContent(t *testing.T) {
	ns := server.New(":8088")
	ns.Start()
	defer ns.Stop()
	files, err := ioutil.ReadDir("../content")
	if err != nil {
		t.Fatalf("can't open content directory: %v", err)
	}
	for _, file := range files {
		contentJSON := "./testdata/" + strings.Replace(file.Name(),"yaml","json",1)
		ep := strings.Split(file.Name(), ".")
		endpoint := "http://localhost:8088/" + ep[0]
	response, err := http.Get(endpoint)
	if err != nil {
		t.Fatalf("could not connect: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("unexpected response status: %v", response.StatusCode)
	}
	defer response.Body.Close()
	got, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("problems reading response body : %v", err)
	}
	want, err := ioutil.ReadFile(contentJSON)
	if err != nil {
		t.Fatalf("problems reading test file: %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(string(want), string(got)))
	}
	}
}
