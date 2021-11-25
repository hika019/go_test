package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	protocol := "tcp"
	port := ":55555"

	tcpAddr, err := net.ResolveTCPAddr(protocol, port)
	checkError(err)

	listner, err := net.ListenTCP(protocol, tcpAddr)
	checkError(err)

	fp, err := os.Create("out.txt")
	checkError(err)
	fp.Close()

	fp, err = os.OpenFile("out.txt", os.O_APPEND|os.O_WRONLY, 0600)

	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		//handleClient(conn)

		defer conn.Close()

		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		fmt.Println("client accept")
		messageBuf := make([]byte, 800)
		messageLen, err := conn.Read(messageBuf)
		checkError(err)

		fmt.Print(messageBuf[:messageLen])

	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	fmt.Println("client accept")
	messageBuf := make([]byte, 800)
	messageLen, err := conn.Read(messageBuf)
	checkError(err)

	fmt.Print(messageBuf[:messageLen])
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
