package main

import (
	"fmt"
	"hex/client"
	"os"
)

func main() {
	var serverAddress, content string
	var listMode bool
	switch len(os.Args) {
	case 1:
		fmt.Fprintf(os.Stderr, "No soup for you")
		os.Exit(1)
	case 2:
		serverAddress = os.Args[1]
		listMode = true
	case 3:
		serverAddress = os.Args[1]
		content = os.Args[2]
	}


	c, err := client.New(serverAddress)
	if err != nil{
		fmt.Printf("can't connect to the server: %v", err)
		os.Exit(1)
	}
	if listMode {
		listOutput, err := c.List()
		if err != nil {
			fmt.Printf("error getting the list: %v", err)
			os.Exit(1)
		}
		fmt.Println(listOutput)
		os.Exit(0)
	}
	c.Start(content)

}