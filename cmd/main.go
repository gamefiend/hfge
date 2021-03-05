package main

import (
	"fmt"
	"hex"
	"log"
)

func main() {
	filename := "content/terrain.yaml"
	hf, err := hex.NewFlowerFromFile(filename)
	if err != nil {
		log.Fatalf("Can't create flower from %q: %v",filename, err)
	}
	for {
		fmt.Println(hf.State())
		fmt.Scanln()
		hf.MoveRandomly()
	}
}

