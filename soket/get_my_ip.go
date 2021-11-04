package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(addrs)
	fmt.Println("------------------")

	for _, addrs := range addrs {
		//fmt.Println(addrs.String())

		if ipnet, ok := addrs.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}
