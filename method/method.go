package main

import "fmt"

func jalanJalan(namaJalan string) string {
	pemberitahuan := "jalan jalan di jalan " + namaJalan

	return pemberitahuan
}

func jumlahDanKali(var1, var2 int16) (jumlah int16, kali int16) {
	jumlah = var1 + var2
	kali = var1 * var2

	return
}

func main() {
	// namaJalan := "Sudirman"
	// runJalanJalan := jalanJalan(namaJalan)

	// fmt.Println(runJalanJalan)

	var var1 int16 = 10
	var var2 int16 = 21

	jumlah, kali := jumlahDanKali(var1, var2)

	fmt.Println("Hasil Jumlah dari : ", var1, " + ", var2, " = ", jumlah)
	fmt.Println("Hasil Kali dari : ", var1, " X ", var2, " = ", kali)

}
