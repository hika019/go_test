package main

import (
	"fmt"
)

func main() {

	var num uint16 = 2040

	hoge := int_to_byte(num)
	fmt.Println(hoge)

	fmt.Println(byte_to_int(hoge))

}
