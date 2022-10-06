package main

import "fmt"

func bukaPintu() string {
	return "Buka Pintu"
}

func tutupPintu() string {
	return "Tutup Pintu"
}

func main() {
	defer fmt.Println(tutupPintu())

	fmt.Println(bukaPintu())
}
