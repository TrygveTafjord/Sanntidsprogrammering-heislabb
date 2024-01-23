package main

import (
	"fmt"
	"net"
)

const LocalIP = "10.100.23.21"
const ServerIp = "10.100.23.129"
const fixedSizeport = "33546"
const serverPort = "20011"

func create_TCP_adress(serverIP string, port string) (*net.TCPAddr, error) {
	addresstring := serverIP + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", addresstring)
	if err != nil {
		return nil, err
	}
	return tcpAddr, nil
}

func TCP_receive(conn *net.TCPConn) (string, error) {
	buffer := make([]byte, 1024)
	size_of_message, err := conn.Read(buffer)
	if err != nil {
		return "There was an error, i didnt read TCP", err
	}
	return string(buffer[:size_of_message]), nil
}

func main() {

	adrString := LocalIP + ":" + serverPort
	listener, err := net.Listen("tcp", adrString)
	if err != nil {
		fmt.Print(("This is wrong x2 wonky"))
		return
	}
	fmt.Print("we started")

	defer listener.Close()

	for {
		fmt.Print("we loop ")
		conn, err := listener.Accept()
		fmt.Print("under accept ")
		if err != nil {
			fmt.Print(("Conn error "))
		}
		fmt.Print("over go ")
		go handleClient(conn)
		fmt.Print("we loop later ")

	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	size_of_message, err := conn.Read(buffer)
	if err != nil {
		fmt.Print("There was an error, i didnt read")
		return
	}
	fmt.Print(buffer[:size_of_message])
}
