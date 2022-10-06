package main

import "fmt"

func main() {
	// conditional statement
	// if false {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// var boolean bool = false
	// if boolean {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// Operator Pembanding
	// var umur uint8 = 34
	// if umur < 34 {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// if umur <= 34 {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// Operator Logika

	// var kanan bool = true
	// var kiri bool = false
	// if kanan && kiri {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// if kanan || kiri {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// if umur == 34 && boolean == true {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// if umur == 34 || boolean == true {
	// 	fmt.Println("kondisi benar")
	// } else {
	// 	fmt.Println("kondisi salah")
	// }

	// Else IF
	// var condition1 uint8 = 10
	// var condition2 uint8 = 15
	// var condition3 uint8 = 20

	// if condition1 >= 13 {
	// 	fmt.Println("kondisi benar 1")
	// } else if condition2 == 15 {
	// 	fmt.Println("kondisi benar 2")
	// } else if condition3 == 34 {
	// 	fmt.Println("kondisi benar 3")
	// } else {
	// 	fmt.Println("kondisi Salah")
	// }

	var condition uint16 = 5000

	if condition >= 13000 {
		fmt.Println("kondisi benar 1")
	} else if condition >= 7000 {
		fmt.Println("kondisi benar 2")
	} else if condition >= 2000 {
		fmt.Println("kondisi benar 3")
	} else {
		fmt.Println("kondisi Salah")
	}

	// var yen uint16
	// yen = condition / 1000
	// if yen >= 15 {
	// 	fmt.Println("Uang anda lebih dari 15 YEN : ", yen, "YEN")
	// } else if yen >= 500 {
	// 	fmt.Println("Uang anda lebih dari 5 YEN : ", yen, "YEN")
	// } else {
	// 	fmt.Println("Uang anda sedikit : ", yen, "YEN")
	// }

	if yen := condition / 1000; yen >= 15 {
		fmt.Println("Uang anda lebih dari 15 YEN", yen, "YEN")
	} else if yen >= 500 {
		fmt.Println("Uang anda lebih dari 5 YEN", yen, "YEN")
	} else {
		fmt.Println("Uang anda sedikit", yen, "YEN")
	}
}
