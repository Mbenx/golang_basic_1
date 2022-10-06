package main

import "fmt"

func main() {
	var barang map[string]string
	barang = map[string]string{}

	barang["pertama"] = "Gelas Kaca"
	barang["kedua"] = "Piring Plastic"

	fmt.Println(barang["pertama"])
	fmt.Println(barang["kedua"])

	code := map[int8]string{}

	code[1] = "Obeng"
	code[12] = "Bor"

	fmt.Println(code[1], code[2])

	nilaiSiswa := map[string]int8{
		"Heru":   80,
		"Fadlan": 85,
		"Kahfy":  90,
	}

	// fmt.Println(nilaiSiswa["Heru"])

	for key, val := range nilaiSiswa {
		fmt.Println(key, " nilainya = ", val)
	}

	var nilaiAkbar, isExist = nilaiSiswa["Akbar"]

	if isExist {
		fmt.Println(nilaiAkbar)
	} else {
		fmt.Println("nilai Akbar tidak di temukan")
	}

	// Slice Map
	alamat := []map[string]string{
		map[string]string{
			"nama":   "heru",
			"alamat": "cawang",
		},
		map[string]string{
			"nama":   "Herlambang",
			"alamat": "Purbalingga",
		},
	}

	// fmt.Println(alamat[0])

	for _, val := range alamat {
		fmt.Println(" Nama : ", val["nama"])
		fmt.Println(" Alamat : ", val["alamat"])
	}

	// Map Slice
	x := map[string][]string{}

	x["buah"] = []string{"jambu", "nanas", "pepaya"}
	x["mobil"] = []string{"audi", "bmw", "merci"}

	fmt.Println(x["buah"])
	fmt.Println(x["buah"][1])

	for key, val := range x {
		fmt.Println(key, " value : ")
		for _, v := range val {
			fmt.Println(" -- ", v)
		}
	}

}
