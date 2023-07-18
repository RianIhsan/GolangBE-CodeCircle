package main

import (
	"fmt"
)

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

func main() {
	fmt.Println(luasPersegiPanjang(20, 20))
	fmt.Println(kelilingPersegiPanjang(20, 20))
	fmt.Println(volumeBalok(20, 20, 10))
}
