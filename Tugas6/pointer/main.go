package main

import (
	"fmt"
	"math"
)

type Lingkaran struct {
	JariJari float64
	Luas     float64
	Keliling float64
}

func (l *Lingkaran) hitungLuasKeliling() {
	l.Luas = math.Pi * l.JariJari * l.JariJari
	l.Keliling = 2 * math.Pi * l.JariJari
}

func main() {
	var lingkaran Lingkaran

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&lingkaran.JariJari)

	lingkaran.hitungLuasKeliling()

	fmt.Printf("Luas lingkaran: %.2f\n", lingkaran.Luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", lingkaran.Keliling)
}
