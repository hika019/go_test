package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s message", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]

	read_file(path)

	/*
		bytes, err := ioutil.ReadFile(path)
		checkError(err)
		fmt.Println(string(bytes))
	*/
}

func read_file(path string) {
	fp, err := os.Open(path)
	checkError(err)

	defer fp.Close()

	buf := make([]byte, 128)
	for {
		bytes, err := fp.Read(buf)
		if bytes == 0 {
			break
		}
		checkError(err)
		fmt.Printf("%d byte\n", bytes)
		fmt.Println(buf)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
