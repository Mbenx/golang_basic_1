package main

import (
	"fmt"
	"math"
)

type Bilangan struct {
	a int8
	b int8
}

func (b Bilangan) tambah() int8 {
	return b.a + b.b
}

type Calculate interface {
	luas() float64
	keliling() float64
}

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	diameter float64
}

func (p Persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p Persegi) keliling() float64 {
	return 4 * p.sisi

}

func (l Lingkaran) luas() float64 {
	jariJari := l.diameter / 2

	return math.Pi * math.Pow(jariJari, 2)
}

func (l Lingkaran) keliling() float64 {
	return math.Pi * l.diameter

}

func main() {
	// versi struct
	// var x = Bilangan{10, 12}
	// fmt.Println(x.tambah())

	// versi interface
	var y Calculate
	y = Persegi{30}
	fmt.Println("===== Persegi ======")
	fmt.Println(y.keliling())
	fmt.Println(y.luas())

	var z Calculate
	z = Lingkaran{30}
	fmt.Println("===== Lingkaran ======")
	fmt.Println(z.keliling())
	fmt.Println(z.luas())

}
