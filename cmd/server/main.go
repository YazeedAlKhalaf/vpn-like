package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	net, err := net.Listen("tcp", ":9195")
	if err != nil {
		fmt.Printf("Failed to listen on port 9195: %v", err)
		os.Exit(1)
	}
	defer net.Close()
	fmt.Println("🚀 Listening on port 9195")

	for {
		conn, err := net.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v", err)
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	length, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading: %s\n", err.Error())
		conn.Write([]byte("Message reading failed :()\n"))
	} else {
		conn.Write([]byte(fmt.Sprintf("Hello %s!", buffer[:length])))
	}

	conn.Close()
}
