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

	protocol := "tcp"
	serverIP := "192.168.11.50"
	serverPort := "55555"
	myIP := "192.168.11.30"
	myPort := 55556

	file_name := os.Args[1]

	fp, err := os.Open(file_name)
	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr(protocol, serverIP+":"+serverPort)
	checkError(err)

	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(myIP)
	myAddr.Port = myPort
	conn, err := net.DialTCP(protocol, myAddr, tcpAddr)
	checkError(err)

	defer conn.Close()

	defer fp.Close()
	sent_binary := make([]byte, 803)

	for {
		bytes, err := fp.Read(sent_binary[:800])
		bytes_size := int_to_byte(uint16(bytes))

		sent_binary[801] = bytes_size[0]
		sent_binary[802] = bytes_size[1]

		if bytes == 0 {
			break
		}
		checkError(err)

		conn.SetDeadline(time.Now().Add(10 * time.Second))
		conn.Write(sent_binary)

		fmt.Printf("%d byte\n", bytes)
		//fmt.Println(sent_binary)
		//fmt.Println(string(sent_binary))
		//fmt.Println(buf)
	}
	fmt.Println("sent the file data")

	conn.SetDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(file_name))
	fmt.Println("Sent the file name")
}
