package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	checkError(err)
	fmt.Println("aaa")
	fmt.Println("GET /HTTP/1.0")

	defer conn.Close()

	readln, err := conn.Read(make([]byte, 1024))
	checkError(err)
	fmt.Println(readln)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
