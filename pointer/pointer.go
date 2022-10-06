package main

import "fmt"

type Product struct {
	id   int
	name string
}

func main() {
	var nomor int16 = 4

	var x int8 = 4

	fmt.Println("Value nomor : ", nomor)
	fmt.Println("Pointer nomor : ", &nomor)

	fmt.Println("Value x : ", x)
	fmt.Println("Pointer x : ", &x)

	var y *int16 = &nomor
	fmt.Println("Pointer y : ", y)
	fmt.Println("value y : ", *y)

	var A [1]Product
	// A[0] = Product{
	// 	id:   1,
	// 	name: "ayam",
	// }
	fmt.Println("value A : ", A[0])

	var B *[1]Product
	// B[0] = Product{
	// 	id:   1,
	// 	name: "ayam",
	// }
	fmt.Println("value B : ", B[0])
}

// Pointer adalah reference atau alamat memori.
// Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
// Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4,
// maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.
