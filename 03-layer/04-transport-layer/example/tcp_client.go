package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to TCP server:", err.Error())
		return
	}
	defer conn.Close()

	message := "Hello, TCP server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message to TCP server:", err.Error())
		return
	}

	fmt.Println("Sent message to TCP server:", message)
}
