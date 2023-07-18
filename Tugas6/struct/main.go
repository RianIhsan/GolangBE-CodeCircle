package main

import "fmt"

type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func main() {

	fmt.Println("NAMA WARNA ADA BIJINYA HARGA")
	dataNanas := buah{
		nama:       "nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}

	fmt.Println(dataNanas)

	dataJeruk := buah{
		nama:       "jeruk",
		warna:      "orange",
		adaBijinya: true,
		harga:      8000,
	}
	fmt.Println(dataJeruk)

	dataSemangka := buah{
		nama:       "semangka",
		warna:      "hijau & merah",
		adaBijinya: true,
		harga:      10000,
	}
	fmt.Println(dataSemangka)

	dataPisang := buah{
		nama:       "pisang",
		warna:      "kuning",
		adaBijinya: false,
		harga:      5000,
	}
	fmt.Println(dataPisang)
}
