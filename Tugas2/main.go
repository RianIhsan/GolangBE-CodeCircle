package main

import (
	"fmt"
)

func calcOperation(operator int, angka1, angka2 int) int {
	switch operator {
	case 1:
		return angka1 + angka2
	case 2:
		return angka1 - angka2
	case 3:
		return angka1 * angka2
	case 4:
		return angka1 / angka2
	default:
		fmt.Println("Operator tidak valid")
		return 0
	}
}

func main() {
	var angka1, angka2 int
	var input int8

	fmt.Println(" INI KALKULATOR SEDERHANA ")
	fmt.Println("1. Pertambahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")
	fmt.Scan(&input)

	fmt.Println("Masukan Angka Awal")
	fmt.Scan(&angka1)
	fmt.Println("Masukan Angka Kedua")
	fmt.Scan(&angka2)

	result := calcOperation(int(input), angka1, angka2)
	fmt.Println("Hasil operasi = ", result)
}
