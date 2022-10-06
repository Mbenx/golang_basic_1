package main

import (
	"fmt"
)

func main() {

	// Restoran -> cek Belanjaan
	// ikan -> cuci -> simpan di frozen
	// daging -> cuci -> simpan di frozen
	// sayur -> kupas -> taruh di box -> simpan di kulkas
	// buah -> cuci -> simpan di kulkas
	// bawang -> kupas -> taruh di box -> simpan di kulkas
	// ! barang di atas -> buang di tempat sampah

	belanja := [15]string{
		"ikan", "daging", "plastik", "sayur", "bungkus",
		"sayur", "ikan", "bawang", "koin", "buah",
		"roti", "gula", "daging", "bawang", "buah",
	}

	for _, barang := range belanja {
		switch barang {
		case "ikan", "daging":
			fmt.Println("menemukan", barang)
			cuci(barang)
		case "sayur", "bawang":
			fmt.Println("menemukan", barang)
			kupas(barang)
		case "buah":
			fmt.Println("menemukan", barang)
			cuci(barang)
		default:
			trashed(barang)
		}
	}

}

func cuci(text string) {
	fmt.Println("mencuci", text)
	if text == "buah" {
		simpanDiKulkas(text)
	} else {
		simpanDiFrozen(text)
	}

}

func kupas(text string) {
	fmt.Println("mengupas", text)
	simpanDiBox(text)
}

func simpanDiBox(text string) {
	fmt.Println("Menyimpan", text, "di Box")
	simpanDiKulkas(text)
}

func simpanDiFrozen(text string) {
	fmt.Println("Menyimpan", text, "di Frozen")
}

func simpanDiKulkas(text string) {
	fmt.Println("Menyimpan", text, "di Kulkas")
}

func trashed(text string) {
	fmt.Println("membuang", text)
}
