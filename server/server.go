package server

import (
	"fmt"
	"hex"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Server struct {
	contentDir string
	address string
	DefaultWebFlowers WebFlowers
	router *mux.Router
}

type WebFlowers map[string]*hex.Flower



type Option func(*Server) error

func WithContentDir(d string) Option {
	return func(s *Server) error {
		content, err := os.ReadDir(d)
		if err != nil {
			return err
		}

		for _, f := range content {
			root := strings.Split(f.Name(), ".")
			endpoint := root[0]
			path := filepath.Join(d, f.Name())
			s.DefaultWebFlowers[endpoint], err = hex.NewFlowerFromFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func WithAddress(address string) Option {
	return func(s *Server) error {
		s.address = address
		return nil
	}
}

func New(opts ...Option) (*Server, error) {
	s := Server {
		DefaultWebFlowers: WebFlowers{},
	}
	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return nil, err
		}
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

