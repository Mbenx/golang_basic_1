package main

import (
	"fmt"
	"strconv"
)

type Buah struct {
	nama  string
	harga int16
}

func (b Buah) hargaBuah() {
	fmt.Println("Harga", b.nama, " = ", b.harga)
}

func (b Buah) beliBuah(name string) string {
	// konversi type data
	var nn = int(b.harga)
	var hargaString = strconv.Itoa(nn)

	var pembelian string = "Buah " + b.nama + " di beli oleh " + name + " dengan harga " + hargaString

	return pembelian
}

func (b Buah) editHargaBuah(hargaBaru int16) string {
	b.harga = hargaBaru
	var hargaString = strconv.Itoa(int(b.harga))
	var hasil string = "Buah " + b.nama + " berubah Harga menjadi " + hargaString

	return hasil

}

func main() {
	var x = Buah{"pepaya", 10000}
	// x.hargaBuah()
	// fmt.Println(x.beliBuah("Heru"))

	fmt.Println(x.editHargaBuah(12500))

}
