package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var x int8 = 10
	var y int16 = 22

	jumlah := int16(x) + y

	fmt.Println(jumlah)

	var a uint16 = 21

	jumlaha := a + uint16(y)
	fmt.Println(jumlaha)

	fmt.Println(reflect.ValueOf(jumlaha).Type())

	// Int to string
	var b = strconv.Itoa(int(y))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Type())

	// string to int
	var c, err = strconv.Atoi("213")
	fmt.Println("err", err)
	fmt.Println(reflect.ValueOf(c))
	fmt.Println(reflect.ValueOf(c).Type())

	// string to float32
	d := "10.23131313"
	var e, err1 = strconv.ParseFloat(d, 32)
	fmt.Println("err", err1)
	fmt.Println(reflect.ValueOf(e))
	fmt.Println(reflect.ValueOf(e).Type())
}
