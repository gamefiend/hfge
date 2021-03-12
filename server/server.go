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


type Server struct {
	contentDir string
	address string
	DefaultWebFlowers WebFlowers
	router *mux.Router
}
type WebFlowers map[string]*hex.Flower

func New(address, contentDir string) (*Server, error) {
	DefaultWebFlowers := WebFlowers{}
	content, err := os.ReadDir(contentDir)
	if err != nil {
		return nil, err
	}

	for _, f := range content {
		root := strings.Split(f.Name(), ".")
		endpoint := root[0]
		path := filepath.Join(contentDir, f.Name())
		DefaultWebFlowers[endpoint], err = hex.NewFlowerFromFile(path)
		if err != nil {
			return nil, err
		}
		//r.HandleFunc(endpoint, handleContent)
		//r.HandleFunc(fmt.Sprintf("%s/{hex}", endpoint), handleContent)
	}
	s := Server {
		address: address,
		contentDir: contentDir,
		DefaultWebFlowers: DefaultWebFlowers,
	}

	s.router = mux.NewRouter()
	s.router.HandleFunc("/{content}/{hex}", s.handleContent)

	return &s, nil
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}

func (s *Server) handleContent(w http.ResponseWriter, r *http.Request){
	// URI: "FLOWER/HEX" ("terrain/10")

	var currentHex int
	var contents string
	url := mux.Vars(r)

	content := url["content"]

	flower := s.DefaultWebFlowers[content]

	// if the flower is blank for some reason or doesn't exist, return a 404 instead of panicking
	if flower == nil {
		http.Error(w, fmt.Sprintf("No value assigned for %s", content), http.StatusNotFound)
		return
	}
	if url["hex"] != ""{
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
