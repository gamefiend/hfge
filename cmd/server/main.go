package main

import (
	"flag"
	"fmt"
	"hex/client"
	"hex/server"
	"log"
)

func main() {
		s, err := server.New(
			server.WithAddress(":8080"),
			server.WithContentDir(*contentPath),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Starting hfge server...")
		log.Fatal(s.Start())
	}
}

