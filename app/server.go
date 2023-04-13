package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6380")
	if err != nil {
		fmt.Println("Failed to bind to port 6380")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Success to bind to port 6380")
	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	if _, err := conn.Read(buffer); err != nil {
		fmt.Println("error reading from client ", err.Error())
	}
	conn.Write([]byte("+PONG\r\n"))

}
