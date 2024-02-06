package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	net, err := net.Dial("tcp", "localhost:9195")
	if err != nil {
		fmt.Printf("Failed to dial: %v\n", err)
		os.Exit(1)
	}
	defer net.Close()

	_, err = net.Write([]byte("Ahmad!"))
	if err != nil {
		fmt.Printf("Failed to write message: %s\n", err.Error())
		return
	}

	buffer := make([]byte, 1024)
	length, err := net.Read(buffer)
	if err != nil {
		fmt.Printf("Failed to read response: %s\n", err.Error())
		return
	}

	fmt.Printf("message received: %s\n", string(buffer[:length]))
}
