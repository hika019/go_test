package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1
	hoge := "aiueo"
	fmt.Println(num)
	fmt.Println(hoge)

	num1 := 123
	var num2 int = 12345
	num3 := 1.235
	var num4 float64 = 1.2345546458

	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num3))
	fmt.Println(reflect.TypeOf(num4))

	flag := 10 > 5
	fmt.Println(flag)
}
