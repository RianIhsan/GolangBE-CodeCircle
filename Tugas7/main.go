package main

import "fmt"

func tanggal(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

func main() {
	kalimat := "golang Backend Development"
	tahun := 2021

	defer func() {
		tanggal(kalimat, tahun)
	}()

	fmt.Println("test")
	fmt.Println("Test Defer")
}
