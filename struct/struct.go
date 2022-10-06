package main

import "fmt"

type Makanan struct {
	sayur Sayur
	Buah
	minuman string
}

type Juice struct {
	Buah
	toping string
}

type Buah struct {
	nama  string
	rasa  string
	harga int16
}

type Sayur struct {
	nama  string
	warna string
	harga int16
}

func main() {
	var makan Makanan
	makan.minuman = "Es Teh"
	makan.sayur.nama = "Bayam"
	makan.sayur.warna = "Hijau"
	makan.Buah = Buah{
		nama: "Apel",
		rasa: "asam",
	}

	// makan.Buah.nama = "Apel"
	// makan.Buah.rasa = "asam"

	fmt.Println(makan)
}
