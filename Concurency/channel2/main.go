package main

import (
	"fmt"
	"time"
)

// global variable
var cuciChannel = make(chan string)
var kupasChannel = make(chan string)
var frozenChannel = make(chan string)
var kulkasChannel = make(chan string)
var boxChannel = make(chan string)
var trashedChannel = make(chan string)

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

	go checkBarang(belanja)
	go cuci()
	go kupas()
	go simpanDiBox()
	go simpanDiFrozen()
	go simpanDiKulkas()
	go trashed()

	time.Sleep(1 * time.Second)

}

func checkBarang(belanja [15]string) {
	for _, barang := range belanja {
		switch barang {
		case "ikan", "daging":
			fmt.Println("menemukan", barang)
			cuciChannel <- barang
		case "sayur", "bawang":
			fmt.Println("menemukan", barang)
			kupasChannel <- barang
		case "buah":
			fmt.Println("menemukan", barang)
			cuciChannel <- barang
		default:
			trashedChannel <- barang
		}
	}
}

func cuci() {
	for itemCuci := range cuciChannel {
		fmt.Println("mencuci", itemCuci)
		if itemCuci == "buah" {
			kulkasChannel <- itemCuci
		} else {
			frozenChannel <- itemCuci
		}
	}

}

func kupas() {
	for itemKupas := range kupasChannel {
		fmt.Println("Mengupas", itemKupas)
		boxChannel <- itemKupas
	}
}

func simpanDiBox() {
	for itemBox := range boxChannel {
		fmt.Println("Menyimpan", itemBox, "di Box")
		kulkasChannel <- itemBox
	}
}

func simpanDiFrozen() {
	for itemFrozen := range frozenChannel {
		fmt.Println("Menyimpan", itemFrozen, "di Frozen")
	}
}

func simpanDiKulkas() {
	for itemKulkas := range kulkasChannel {
		fmt.Println("Menyimpan", itemKulkas, "di Kulkas")
	}
}

func trashed() {
	for itemTrashed := range trashedChannel {
		fmt.Println("Membuang", itemTrashed)
	}
}
