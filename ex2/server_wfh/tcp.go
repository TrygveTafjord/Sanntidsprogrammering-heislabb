package main

import (
	"fmt"
	"net"
)

const LocalIP = "10.100.23.17"
const ServerIp = "10.100.23.129"
const fixedSizeport = "34933" 
const serverPort = "20007"

func create_TCP_adress(serverIP string, port string) (*net.TCPAddr, error) {
	addresstring := serverIP + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", addresstring)
	if err != nil {
		return nil, err
	}
	return tcpAddr, nil
}

// func TCP_send(conn *net.UDPConn, stringToSend string) {
// 	_, err := conn.Write([]byte(stringToSend))
// 	if err != nil {
// 		fmt.Print("There has been an error in UDP_send")
// 	}
// }

func TCP_receive(conn *net.TCPConn) (string, error) {
	buffer := make([]byte, 1024)
	size_of_message, err := conn.Read(buffer)
	if err != nil {
		return "There was an error, i didnt read", err
	}
	return string(buffer[:size_of_message]), nil
}

func main() {
    adrString := ServerIp + ":" + fixedSizeport
	//  create_TCP_adress(ServerIp, fixedSizeport)
	// if err != nil {
	// 	fmt.Print(("This is wrong"))
	// }

	listener, err := net.ResolveTCPAddr("tcp", adrString)
	if err != nil {
		fmt.Print(("This is wrong"))
	}
	conn, err := net.ListenTCP("udp", listener)
	if err != nil {
		fmt.Print(("This is wrong x2"))
	}
	defer conn.Close()


	// conn, err := net.DialUDP("udp4", nil, udpAddr)
	// if err != nil {
	// 	fmt.Print("ERROR in dial up")
	// }
	// defer conn.Close()

	returned_message, _ := TCP_receive(conn)
	fmt.Printf("Vi fikk: %s", returned_message)

}
