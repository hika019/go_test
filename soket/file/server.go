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
	tmp_file_name := "tmp.txt"

	fp, err := os.OpenFile(tmp_file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		messageLen, err := conn.Read(messageBuf)
		checkError(err)
		if messageLen == 0 {
			break
		}
		//ファイルに書き込み
		fp.Write(messageBuf[:messageLen])

	}
	fmt.Println("Downloaded file data")

	//ファイル名の変更
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	messageLen, err := conn.Read(messageBuf)
	checkError(err)
	fp.Close()

	file_name := string(messageBuf[:messageLen-16])
	err = os.Rename(tmp_file_name, file_name)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
