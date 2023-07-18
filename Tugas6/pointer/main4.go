package main

import "fmt"

type Film struct {
	Judul    string
	Durasi   string
	Category string
	Release  string
}

func addToFilm(judul, durasi, category, release string, dataFilm *[]Film) {
	film := Film{
		Judul:    judul,
		Durasi:   durasi,
		Category: category,
		Release:  release,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	var dataFilm = []Film{}

	addToFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	addToFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	addToFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	addToFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. Judul: %s, Durasi: %s, Kategori: %s, Tahun Rilis: %s\n", i+1, item.Judul, item.Durasi, item.Category, item.Release)
	}
}
