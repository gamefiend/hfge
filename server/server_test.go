package server_test

import (
	"encoding/json"
	"fmt"
	"hex/server"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestServerFailsToInitializeWithWrongFilePath(t *testing.T){
	_, err := server.New(":8083", "./randomjunk")
	if err == nil {
		t.Errorf("Server should return error when given erroneous filepath")
	}

}

func TestBlankFlower(t *testing.T) {
	// this is the wrong place to get data from, so the webflower will be blank
	ns, err := server.New(":8087", "../testdata")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		t.Fatal(ns.Start())
	}()
	time.Sleep(1 * time.Second)

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
	ns, err := server.New(":8088", "../content")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		t.Fatal(ns.Start())
	}()
	time.Sleep(1 * time.Second)
	defer ns.Stop()
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

func inNeighborList(hex int, neighbors []int) bool {
	for _, n := range neighbors {
		if n == hex {
			return true
		}
	}
	return false
}