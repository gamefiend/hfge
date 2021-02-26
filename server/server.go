package server

import (
	"bytes"
	"encoding/json"
	"github.com/ghodss/yaml"
	"hex"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

func (s *Server) Start(){
	content, err := ioutil.ReadDir("../content")
	if err != nil {
		log.Print(err)
	}
	for _, f := range content {
		root := strings.Split(f.Name(), ".")
		endpoint := "/" + root[0]
		DefaultWebFlowers[endpoint], err = hex.NewFlowerFromFile(f.Name())
		if err != nil {
			log.Print(err)
		}
		log.Print(DefaultWebFlowers[endpoint])
		http.Handle(endpoint, http.HandlerFunc(handleContent))
	}
	go http.ListenAndServe(s.address, nil)
}

func handleContent(w http.ResponseWriter, r *http.Request){
	content := "../content" + r.URL.RequestURI() + ".yaml"
	yamlContent, err := ioutil.ReadFile(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonContent, err := yaml.YAMLToJSON(yamlContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonOutput, err := prettyJSON(jsonContent)
	if err != nil {
		log.Print(err)
	}
	w.Write(jsonOutput)
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
