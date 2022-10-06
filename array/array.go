package main

import "fmt"

func main() {
	// array
	// var a [2]string
	// a[0] = "hello"
	// a[1] = "world"

	// fmt.Println(a)
	// fmt.Println(a[1])
	// fmt.Println(a[0])

	b := [7]int{10, 9, 8, 7, 6, 5, 4}
	fmt.Println(b)

	// // slice
	// var sliceNilai []int = b[1:6] // [A:B]ambil element ke A hinga sebelum element ke B
	// fmt.Println(sliceNilai)

	var sliceNilai1 []int = b[:6]
	// var sliceNilai1 [6]int = b[:6]
	fmt.Println(sliceNilai1)

	// var sliceNilai2 []int = b[2:]
	// fmt.Println(sliceNilai2)

	// array multi dimensi
	// var satu = [2][3]int{
	// 	[3]int{1, 2, 3},
	// 	[3]int{40, 123, 131}}

	var dua = [2][3]int{{1, 2, 3}, {40, 123, 131}}

	// fmt.Println(satu)
	fmt.Println(dua)

	// define array, tanpa declare index of array
	// var aa = [...]int{1, 2, 3, 7, 8}
	// fmt.Println(len(aa))

	// for i := 0; i < len(aa); i++ {
	// 	fmt.Println("Nilai dari aa ke i : ", aa[i])
	// }

	// for a := range aa {
	// 	fmt.Println("Nilai dari aa ke i : ", a)
	// }

	// var numbers = [...]string{"satu", "dua", "tiga"}
	// for i, number := range numbers {
	// 	fmt.Println("Show the Number  of : ", i, number)
	// }

}
