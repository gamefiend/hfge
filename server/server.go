package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"hex"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
	content, err := os.ReadDir(ContentDir)
	if err != nil {
		return err
	}
	r := mux.NewRouter()
	for _, f := range content {
		root := strings.Split(f.Name(), ".")
		endpoint := "/" + root[0]
		path := filepath.Join(ContentDir, f.Name())
		DefaultWebFlowers[endpoint], err = hex.NewFlowerFromFile(path)
		if err != nil {
			return err
		}
		r.HandleFunc(endpoint, handleContent)
		r.HandleFunc(fmt.Sprintf("%s/{hex}", endpoint), handleContent)
	}
	go http.ListenAndServe(s.address, r)
	return nil
}

func handleContent(w http.ResponseWriter, r *http.Request){
	// URI: "FLOWER/HEX" ("terrain/10")
	var currentHex int
	var contents string
	url := mux.Vars(r)
	endpoint := getRootSubTree(r.URL.RequestURI())
	flower := DefaultWebFlowers[endpoint]
	if url["hex"] != "" {
		ch, _ := strconv.Atoi(url["hex"])
		flower.SetHex(ch)
		flower.MoveRandomly()
	}
	currentHex = flower.CurrentHex()
	contents = flower.State()
	jsonOutput := `{
  "current_hex": %d,
  "content": %q
}`
	fmt.Fprintf(w, jsonOutput, currentHex, contents)
}

func (s *Server) Stop(){

}

func getRootSubTree(URL string) string {
	root := strings.Split(URL, "/")
	return fmt.Sprintf("/%s",root[1])
}

func prettyJSON(b []byte) ([]byte, error){
	var out bytes.Buffer
	//get proper formatting so we return something readable
	err := json.Indent(&out, b, "", "  ")
	// append a newline
	return append(out.Bytes(), "\n"...), err
}
