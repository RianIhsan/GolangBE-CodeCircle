package main

import (
	"fmt"
)

type Satuan struct {
	Panjang int32
	Lebar   int32
	Tinggi  int32
}

func main() {
	satuan := Satuan{
		Panjang: 7,
		Lebar:   4,
		Tinggi:  6,
	}

	volumeBalok := satuan.Panjang * satuan.Lebar * satuan.Tinggi

	fmt.Println("Panjang =", satuan.Panjang)
	fmt.Println("Lebar =", satuan.Lebar)
	fmt.Println("Tinggi =", satuan.Tinggi)
	fmt.Printf("Volume Balok = %d", volumeBalok)
}
