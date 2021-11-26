package main

import (
	"crypto/rand"
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

	}
}

func handleClient(conn net.Conn) {

	addr, ok := conn.RemoteAddr().(*net.TCPAddr)
	if !ok {
		return
	}

	fmt.Println(addr.IP.String())

	defer conn.Close()
	messageBuf := make([]byte, 803)
	tmp_file_name := "tmp-" + MakeRandomStr() + ".txt"

	fp, err := os.OpenFile(tmp_file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		messageLen, err := conn.Read(messageBuf)
		checkError(err)
		fmt.Printf("%d byte\n", messageLen)
		//fmt.Println(messageBuf[:messageLen])
		if messageLen != 803 {
			fmt.Println("Downloaded file data")
			fp.Close()

			file_name := string(messageBuf[:messageLen])
			err = os.Rename(tmp_file_name, file_name)
			checkError(err)
			fmt.Println("Change file name: ", file_name)
			break
		}
		//ファイルに書き込み

		data_size_byte := make([]byte, 2)
		data_size_byte[0] = messageBuf[801]
		data_size_byte[1] = messageBuf[802]
		data_size := byte_to_int(data_size_byte)
		fmt.Printf("%d byte\n", data_size)

		fmt.Fprintf(fp, "%s", string(messageBuf[:data_size]))

	}

}

func MakeRandomStr() string {
	b := make([]byte, 15)
	_, err := rand.Read(b)
	checkError(err)

	var str string
	for _, v := range b {
		str += string(v%byte(94) + 33)
	}
	return str
}
