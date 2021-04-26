package server_test

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"hex/server"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)
// Tests that test a faulty server should initialize their own server while
// those that assume a correctly configured server use the server spun up
// in the TestMain function.

func TestServerFailsToInitializeWithWrongFilePath(t *testing.T){

	_, err := server.New(
		server.WithAddress(":8083"),
		server.WithContentDir("./randomjunk"),
		)
	if err == nil {
		t.Errorf("Server should return error when given erroneous filepath")
	}

}

func TestBlankFlower(t *testing.T) {
	// this is the wrong place to get data from, so the webflower will be blank
	ns, err := server.New(
		server.WithAddress(":8087"),
		server.WithContentDir("../testdata"),
		)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		t.Fatal(ns.Start())
	}()
	waitForServer(":8087")
	response, err := http.Get("http://localhost:8087/terrain/1")
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("expected %d, got error code %d", http.StatusNotFound, response.StatusCode)
	}

}

func TestMove(t *testing.T) {
	var fr struct {
		CurrentHex int `json:"current_hex"`
		Content string `json:"content"`
	}
	tcs := []struct {
		name string
		URI  string
		wantOneOf []int
	}{
		{
			name: "terrain_4",
			URI:  "/terrain/4",
			wantOneOf: []int{0, 9, 7, 2, 0, 0},
		},
		{
			name: "terrain_start",
			URI:  "/terrain/5",
			wantOneOf: []int{7, 10, 8, 3, 1, 2},
		},
		{
			name: "terrain_start",
			URI:  "/terrain/19",
			wantOneOf: []int{0, 0, 0, 18, 15, 17},
		},
		{
			name: "weather_4",
			URI:  "/weather/4",
			wantOneOf: []int{0, 9, 7, 2, 0, 0},
		},
		{
			name: "weather_start",
			URI:  "/weather/10",
			wantOneOf: []int{12, 15, 13, 8, 5, 7},
		},
		{
			name: "weather_start",
			URI:  "/weather/7",
			wantOneOf: []int{9, 12, 10, 5, 2, 4},
		},
	}
	for _, tc := range tcs {
		response, err := http.Get(fmt.Sprintf("http://localhost:8088/%s", tc.URI))
		if err != nil {
			t.Fatalf("could not connect: %v", err)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatalf("unexpected response status for %v: %v", response.Request.URL.Path, response.StatusCode)
		}
		defer response.Body.Close()
		got, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatalf("problems reading response body : %v", err)
		}
		err = json.Unmarshal(got, &fr)
		if err != nil{
			t.Fatalf("problems parsing json response: %v", err)
		}
		if !inNeighborList(fr.CurrentHex, tc.wantOneOf) {
			t.Errorf("Should have moved to one of : %v , got: %v", tc.wantOneOf, fr.CurrentHex)
		}
	}
}

func TestListDisplaysCorrectly(t *testing.T){
	want := "terrain\nweather\n"
	response, err := http.Get("http://localhost:8088/list")
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		r, _ := ioutil.ReadAll(response.Body)
		t.Fatalf("unexpected response status for %v: %v %v", response.Request.URL.Path, string(r), response.StatusCode)
	}
	got, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, string(got)){
		t.Error(cmp.Diff(want, string(got)))
	}
}

func inNeighborList(hex int, neighbors []int) bool {
	for _, n := range neighbors {
		if n == hex {
			return true
		}
	}
	return false
}

func TestMain(m *testing.M){
	// we don't want to keep spinning up a server and having to wait for it each time,
	// so we spin up a server before tests and wait for it once.
	port := ":8088"
	content := "../content"
	ns, err := server.New(
		server.WithAddress(port),
		server.WithContentDir(content),)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(ns.Start())
	}()
	err = waitForServer(port)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func waitForServer(port string) error{
	timeout := 1 * time.Second
	deadline := time.Now().Add(1 * time.Second)
	n := net.Dialer{Timeout: timeout,
	Deadline: deadline}
	address := "localhost" + port
	conn, err := n.Dial("tcp", address)
	if err !=nil {
		return err
	}
	defer conn.Close()
	return nil
}

