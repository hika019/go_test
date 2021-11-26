package main

import (
	"fmt"
	"os"
)

func int_to_byte(i uint16) []byte {

	byte_data := make([]byte, 2)

	byte_data[0] = uint8(i % 256)
	byte_data[1] = uint8((i / 256) % 256)
	return byte_data[:]
}

func byte_to_int(byte_data []byte) uint16 {
	num := int(byte_data[0])
	num += int(byte_data[1]) * 256
	return uint16(num)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal: error: ", err.Error())
		os.Exit(1)
	}
}
