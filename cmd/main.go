package main

import (
	"flag"
	"fmt"
	"hex"
	"hex/server"
	"log"
)

func main() {
	client := flag.Bool("client", false, "activates hfge client mode")
	content := flag.String("contentPath", "./content", "specifies location of content directory")
	flag.Parse()
	if *client {
		fmt.Println("hfge client")
		filename := "content/terrain.yaml"
		hf, err := hex.NewFlowerFromFile(filename)
		if err != nil {
			log.Fatalf("Can't create flower from %q: %v", filename, err)
		}
		for {
			fmt.Println(hf.State())
			fmt.Scanln()
			hf.MoveRandomly()
		}
	} else {
		s, err := server.New(":8080", *content)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("Starting hfge server...")
		 log.Fatal(s.Start())

	}
}

