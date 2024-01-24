package main

import (
	"fmt"
	"net"
)

const LocalIP = "10.100.23.17"
const ServerIp = "10.100.23.129"
const group8IP = "10.100.23.18"
const group26IP = "10.100.23.36"
const serverPort = "20007"

func create_UDP_adress(serverIP string, port string) (*net.UDPAddr, error) {
	addresstring := serverIP + ":" + port
	udpAddr, err := net.ResolveUDPAddr("udp4", addresstring)
	if err != nil {
		return nil, err
	}
	return udpAddr, nil
}

func UDP_send(conn *net.UDPConn, stringToSend string) {
	_, err := conn.Write([]byte(stringToSend))
	if err != nil {
		fmt.Print("There has been an error in UDP_send")
	}
}

func UDP_receive(conn *net.UDPConn) (string, error) {
	buffer := make([]byte, 1024)
	size_of_message, err := conn.Read(buffer)
	if err != nil {
		return "There was an error, i didnt read", err
	}
	return string(buffer[:size_of_message]), nil
}

func main() {

	listener, err := net.ResolveUDPAddr("udp", ":30000")
	if err != nil {
		fmt.Print(("This is wrong"))
	}
	conn1, err := net.ListenUDP("udp", listener)
	if err != nil {
		fmt.Print(("This is wrong x2"))
	}
	defer conn1.Close()

	udpAddr, err := create_UDP_adress(ServerIp, "20007")
	if err != nil {
		fmt.Print("ERROR")
	}

	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		fmt.Print("ERROR in dial up")
	}
	defer conn.Close()

	UDP_send(conn, "Dette funker g7")
	returned_message, _ := UDP_receive(conn1)
	fmt.Printf("Vi fikk: %s", returned_message)

}
