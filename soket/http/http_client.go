package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	conn, err := http.Get("http://www.google.com/")
	checkError(err)

	defer conn.Body.Close()

	body, err := ioutil.ReadAll(conn.Body)
	checkError(err)
	fmt.Println(string(body))

}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "fatal: error: &s", err.Error())
		os.Exit(1)
	}
}
