package main

import "fmt"

func konversiNilai(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else if nilai >= 60 && nilai < 70 {
		return "C"
	} else if nilai >= 50 && nilai < 60 {
		return "D"
	} else if nilai < 50 {
		return "E"
	} else {
		return "wrong option"
	}
}

func main() {
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println(konversiNilai(nilaiJohn))
	fmt.Println(konversiNilai(nilaiDoe))
}
