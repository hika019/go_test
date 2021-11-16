package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s message", os.Args[0])
		os.Exit(1)
	}

	message := os.Args[1]

	protocol := "tcp"
	serverIP := "192.168.11.30"
	serverPort := "55555"
	//myIP := get_my_ip()
	myIP := "192.168.11.50"
	myPort := 55556

	fmt.Println("init setting")
	fmt.Println(serverIP)

	tcpAddr, err := net.ResolveTCPAddr(protocol, serverIP+":"+serverPort)
	checkError(err)

	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(myIP)
	myAddr.Port = myPort
	conn, err := net.DialTCP(protocol, myAddr, tcpAddr)
	checkError(err)

	fmt.Println("TCPconnect")

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))

	readBuf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readlen, err := conn.Read(readBuf)
	checkError(err)

	fmt.Println("server: " + string(readBuf[:readlen]))
}

func checkError(err error) {
	if err != nil {
		fmt.Print("fatal: error: &s", err.Error())
		os.Exit(1)
	}
}

func get_my_ip() string {
	addrs, err := net.InterfaceAddrs()
	checkError(err)

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {

			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}

	fmt.Fprint(os.Stderr, "failed get IP address")
	os.Exit(1)
	return ""
}
