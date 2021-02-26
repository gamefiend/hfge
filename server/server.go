package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hex"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const ContentDir = "../content"

type Server struct {
	address string
}
type WebFlowers map[string]*hex.Flower
var DefaultWebFlowers = WebFlowers{}

func New(address string) *Server {
	return &Server {
		address: address,
	}
}

func (s *Server) Start() error {
	fmt.Println("entering Start")
	content, err := os.ReadDir(ContentDir)
	if err != nil {
		return err
	}
	for _, f := range content {
		root := strings.Split(f.Name(), ".")
		endpoint := "/" + root[0]
		path := filepath.Join(ContentDir, f.Name())
		fmt.Println("filename", path)
		DefaultWebFlowers[endpoint], err = hex.NewFlowerFromFile(path)
		if err != nil {
			return err
		}
		log.Print(DefaultWebFlowers[endpoint])
		http.Handle(endpoint, http.HandlerFunc(handleContent))
	}
	go http.ListenAndServe(s.address, nil)
	return nil
}

func handleContent(w http.ResponseWriter, r *http.Request){
	flower := DefaultWebFlowers[r.URL.RequestURI()]
	fmt.Println("Current hex", flower.CurrentHex())
	currentHex := flower.CurrentHex()
	contents := flower.State()
	jsonOutput := `{
  "current_hex": %d,
  "content": %q
}`
	fmt.Fprintf(w, jsonOutput, currentHex, contents)
}

func (s *Server) Stop(){

}

func prettyJSON(b []byte) ([]byte, error){
	var out bytes.Buffer
	//get proper formatting so we return something readable
	err := json.Indent(&out, b, "", "  ")
	// append a newline
	return append(out.Bytes(), "\n"...), err
}
