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

	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
		/*
			defer conn.Close()

			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			fmt.Println("client accept")
			messageBuf := make([]byte, 800)
			messageLen, err := conn.Read(messageBuf)
			checkError(err)

			fmt.Print(messageBuf[:messageLen])
		*/
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	messageBuf := make([]byte, 800)
	file_name := "tmp.txt"

	flag := true
	var fp *os.File
	var err error

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		conn.Read(messageBuf)
		if flag == true {
			file_name = string(messageBuf)
			fp, err = os.Create(file_name)
			checkError(err)
			fp.Close()
			fmt.Println("get the file name")
			flag = false
		} else {
			fp, err = os.OpenFile(file_name, os.O_APPEND|os.O_WRONLY, 0600)
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			fmt.Println("client accept")

			messageLen, err := conn.Read(messageBuf)
			checkError(err)
			if messageLen == 0 {
				break
			}

			fmt.Print(messageBuf[:messageLen])
			fmt.Fprintf(fp, "%s", string(messageBuf[:messageLen]))
			fp.Close()
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
