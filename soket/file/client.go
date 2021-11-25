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
	/*
		//ファイル名の転送
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		conn.Write([]byte(path))

		//ファイル名の送信ができたか
		readBuff := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		readlen, err := conn.Read(readBuff)
		checkError(err)
		if int(readlen) != 0 {
			fmt.Fprintln(os.Stderr, "fatal: error: can't sent file name")
			os.Exit(1)
		}
	*/

	defer fp.Close()
	sent_binary := make([]byte, 800)

	conn.SetDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(file_name))
	fmt.Println("Sent the file name")

	for {
		bytes, err := fp.Read(sent_binary)
		if bytes == 0 {
			break
		}
		checkError(err)

		conn.SetDeadline(time.Now().Add(10 * time.Second))
		conn.Write(sent_binary)

		fmt.Printf("%d byte\n", bytes)
		fmt.Println(string(sent_binary))
		//fmt.Println(buf)

	}

	/*
		bytes, err := ioutil.ReadFile(path)
		checkError(err)
		fmt.Println(string(bytes))
	*/
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
