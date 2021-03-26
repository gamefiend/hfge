package main

import (
	"flag"
	"fmt"
	"hex/client"
	"os"
)

func main() {
	var serverAddress, content string
	var listMode bool
	switch len(os.Args) {
	case 0:
		fmt.Fprintf(os.Stderr, "No soup for you")
		os.Exit(1)
	case 1:
		serverAddress = os.Args[1]
		listMode = true
	case 2:
		serverAddress = os.Args[1]
		content = os.Args[2]
	}

	flag.Parse()
	c := client.New(serverAddress)
	if listMode {
		c.List()
		os.Exit(0)
	}
	c.Start(content)
}