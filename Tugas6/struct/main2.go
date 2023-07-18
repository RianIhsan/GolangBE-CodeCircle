package main

import "fmt"

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func luasSegitiga(s segitiga) int {
	return (s.alas * s.tinggi) / 2
}

func luasPersegi(p persegi) int {
	return (p.sisi * p.sisi)
}

func luasPersegiPanjang(pp persegiPanjang) int {
	return (pp.panjang * pp.lebar)
}

func main() {
	hSegitiga := segitiga{alas: 4, tinggi: 3}
	hPersegi := persegi{sisi: 5}
	hPersegiPanjang := persegiPanjang{panjang: 5, lebar: 3}

	fmt.Println("luas segitiga : ", luasSegitiga(hSegitiga))
	fmt.Println("luas persegi : ", luasPersegi(hPersegi))
	fmt.Println("luas persegi panjang :", luasPersegiPanjang(hPersegiPanjang))
}
