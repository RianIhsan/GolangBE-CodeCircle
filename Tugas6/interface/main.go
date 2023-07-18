package main

import (
	"fmt"
	"math"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}
type persegiPanjang struct {
	panjang, lebar int
}
type tabung struct {
	jariJari, tinggi float64
}
type balok struct {
	panjang, lebar, tinggi int
}
type hitungBangunDatar interface {
	luas() int
	keliling() int
}
type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() float64 {
	return (float64(s.alas*s.tinggi) / 2)
}

func (p persegiPanjang) keliling() float64 {
	return (float64(2*p.panjang + p.lebar))
}

func (b balok) volume() float64 {
	return (float64(b.panjang * b.lebar * b.tinggi))
}

func (bKeliling tabung) keliling() float64 {
	keliling_alas := 2 * math.Pi * bKeliling.jariJari
	keliling_selimut := 2 * math.Pi * bKeliling.jariJari * bKeliling.tinggi
	return (float64(keliling_alas + keliling_selimut))
}

func main() {
	hSss := &segitigaSamaSisi{alas: 15, tinggi: 3}
	hPpp := &persegiPanjang{panjang: 15, lebar: 3}

	hBalok := &balok{panjang: 2, lebar: 5, tinggi: 20}
	hKeliling := &tabung{jariJari: 20, tinggi: 5}

	fmt.Println("segitiga sama sisi :", hSss.luas())
	fmt.Println("keliling persegi panjang :", hPpp.keliling())
	fmt.Println("")
	fmt.Println("tabung : ", hBalok.volume())
	fmt.Println("keliling :", hKeliling.keliling())
}
