package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "www.google.com:80")
	checkError(err)

	request_text := "GET /HTTP/1.0\r\n\r\n"
	request_byte := []byte(request_text)

	fmt.Println(request_byte)

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write(request_byte)

	fmt.Println("send message")

	readBuf := make([]byte, 1024)
	readln, err := conn.Read(readBuf)
	checkError(err)

	fmt.Println(string(readBuf[:readln]))

}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
