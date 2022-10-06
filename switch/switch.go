package main

import "fmt"

func main() {
	switchVar := "Biru"
	// fmt.Println(switchVar)

	switch switchVar {
	case "Biru":
		fmt.Println("Case Pertama")
	case "Kuning":
		fmt.Println("Case Keduan")
	case "yellow", "red", "brown":
		fmt.Println("Case Ketiga")
	default:
		fmt.Println("Case Default")
	}

	switchVar2 := 15
	switch {
	case switchVar2 == 8:
		fmt.Println("Case Pertama")
		fallthrough
	case switchVar2 == 13:
		fmt.Println("Case kedua")

	case switchVar2 == 15:
		fmt.Println("Case ketiga")

	default:
		fmt.Println("Case Default")
	}

	// untuk analogi
	auth := "ADMIN"
	switch auth {
	case "Admin":
		// masuk halaman admin
	case "staff":
		// masuk halaman staff
	case "manager", "general manager":
		// halaman leader
	default:
		// halaman tamu
	}
}
