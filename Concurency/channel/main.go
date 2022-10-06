package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var print = func(name string) {
		// fmt.Println("Hello", name)
		var data = "hello " + name
		messages <- data
	}

	go print("heru")
	go print("sopan")
	go print("ali")

	var data1 = <-messages
	fmt.Println(data1)

	var data2 = <-messages
	fmt.Println(data2)

	var data3 = <-messages
	fmt.Println(data3)

	// var data4 = <-messages
	// fmt.Println(data4)

}
