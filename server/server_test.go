package server_test

import (
	"encoding/json"
	"fmt"
	"hex/server"
	"io"
	"net/http"
	"testing"
)


type FlowerResponse struct {
	CurrentHex int `json:"current_hex"`
	Content string `json:"content"`
}

func TestMove(t *testing.T) {
	var fr FlowerResponse
	ns := server.New(":8088")
	ns.Start()
	defer ns.Stop()
	tcs := []struct {
		name string
		URI  string
		want []int
	}{
		{
			name: "terrain_4",
			URI:  "/terrain/4",
			want: []int{0, 9, 7, 2, 0, 0},
		},
		{
			name: "terrain_start",
			URI:  "/terrain/5",
			want: []int{7, 10, 8, 3, 1, 2},
		},
		{
			name: "terrain_start",
			URI:  "/terrain/19",
			want: []int{0, 0, 0, 18, 15, 17},
		},
		{
			name: "weather_4",
			URI:  "/weather/4",
			want: []int{0, 9, 7, 2, 0, 0},
		},
		{
			name: "weather_start",
			URI:  "/weather/10",
			want: []int{12, 15, 13, 8, 5, 7},
		},
		{
			name: "weather_start",
			URI:  "/weather/7",
			want: []int{9, 12, 10, 5, 2, 4},
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
		matches := func() bool {
			for _, n := range tc.want {
				if n == fr.CurrentHex {
					return true
				}
			}
			return false
		}
		// want to make sure that the result is on the actual list of neighbors
		if !matches() {
			t.Errorf("Should have moved to one of : %v , got: %v", tc.want, fr.CurrentHex)

		}
	}
}

