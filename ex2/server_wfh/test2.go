package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Server IP and ports
	serverIP := "10.100.23.129" // Replace with actual server IP
	fixedSizePort := "34933"
	delimiterPort := "33546"

	// Connect to server for fixed size messages
	fixedSizeConn, err := net.Dial("tcp", serverIP+":"+fixedSizePort)
	if err != nil {
		fmt.Println("Error connecting to fixed size port:", err)
		os.Exit(1)
	}
	defer fixedSizeConn.Close()

	// Connect to server for delimited messages
	delimiterConn, err := net.Dial("tcp", serverIP+":"+delimiterPort)
	if err != nil {
		fmt.Println("Error connecting to delimiter port:", err)
		os.Exit(1)
	}
	defer delimiterConn.Close()

	// Send a message to fixed size port
	fixedSizeMsg := "Hello from TCP fixed size client!"
	fixedSizeConn.Write([]byte(fixedSizeMsg))

	// Send a message to delimiter port
	delimiterMsg := "Hello from TCP delimiter client!\000"
	delimiterConn.Write([]byte(delimiterMsg))

	// Read response from fixed size port
	fixedSizeResponse := make([]byte, 1024)
	fixedSizeConn.Read(fixedSizeResponse)
	fmt.Println("Fixed size response:", string(fixedSizeResponse))

	// Read response from delimiter port
	delimiterResponse, _ := bufio.NewReader(delimiterConn).ReadString('\000')
	fmt.Println("Delimiter response:", strings.TrimRight(delimiterResponse, "\000"))
}
