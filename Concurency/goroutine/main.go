package main

import (
	"fmt"
	"time"
)

func print(a int, message string) {
	for i := 0; i < a; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// runtime.GOMAXPROCS(2)

	// go print(5, "cek")
	go print(5, "abc")

	// var input string
	// fmt.Scanln(&input)

	time.Sleep(1 * time.Second)

}
