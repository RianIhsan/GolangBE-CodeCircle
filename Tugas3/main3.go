package main

import (
	"fmt"
)

func hitungPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func hitungLuasSegitiga(alas, tinggi int) float32 {
	return 0.5 * float32(alas) * float32(tinggi)
}

func main() {
	var panjangPersegiPanjang, lebarPersegiPanjang int
	var alasSegitiga, tinggiSegitiga int

	fmt.Println("Masukkan panjang persegi panjang:")
	fmt.Scan(&panjangPersegiPanjang)
	fmt.Println("Masukkan lebar persegi panjang:")
	fmt.Scan(&lebarPersegiPanjang)

	hasilPersegiPanjang := hitungPersegiPanjang(panjangPersegiPanjang, lebarPersegiPanjang)
	fmt.Println("Hasil persegi panjang:", hasilPersegiPanjang)

	fmt.Println("Masukkan alas segitiga:")
	fmt.Scan(&alasSegitiga)
	fmt.Println("Masukkan tinggi segitiga:")
	fmt.Scan(&tinggiSegitiga)

	hasilLuasSegitiga := hitungLuasSegitiga(alasSegitiga, tinggiSegitiga)
	fmt.Println("Hasil luas segitiga:", hasilLuasSegitiga)
}
