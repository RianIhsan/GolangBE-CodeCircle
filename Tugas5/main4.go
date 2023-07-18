package main

import (
	"fmt"
)

func main() {
	var dataFilm []map[string]string

	addDataFilm := func(title, duration, genre, year string) {
		film := map[string]string{
			"title":    title,
			"duration": duration,
			"genre":    genre,
			"year":     year,
		}

		dataFilm = append(dataFilm, film)
	}

	addDataFilm("Naruto The Movie 2", "2 Jam", "Action", "2023")
	addDataFilm("One Piece Arch", "2,5 Jam", "Adventure", "2023")
	addDataFilm("Jujutsu Kaise", "2 Jam", "Action", "2022")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
