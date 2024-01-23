package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "10.100.23.129:20011")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	// Read welcome message
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		os.Exit(1)
	}
	fmt.Println("Dette kjører")

	welcomeMessage := string(buffer[:n])
	fmt.Println("Welcome message:", welcomeMessage)

	// Send and receive messages
	for {
		message := "Hello, server!"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing:", err)
			os.Exit(1)
		}

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			os.Exit(1)
		}

		response := string(buffer[:n])
		fmt.Println("Server response:", response)
	}
}
