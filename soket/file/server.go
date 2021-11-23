package main

import (
	"fmt"
	"net"
	"os"
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

		defer conn.Close()

		Buf := make([]byte, 800)
		Buf_len, err := conn.Read(Buf)
		checkError(err)

		if Buf_len == 0 {
			break
		}

		fmt.Fprintf(fp, "%s", string(Buf[:Buf_len]))

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
