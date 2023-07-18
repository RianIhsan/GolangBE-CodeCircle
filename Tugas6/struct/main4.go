package main

import "fmt"

type movie struct {
	title, genre   string
	duration, year int
}

func addMovie(title, genre string, duration, year int, dataFilm *[]movie) {
	film := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}

	*dataFilm = append(*dataFilm, film)
}

func displayMovies(dataFilm *[]movie) {
	for i, film := range *dataFilm {
		fmt.Printf("%d. title : %s \n", i+1, film.title)
		fmt.Printf("duration : %d \n", film.duration)
		fmt.Printf("genre : %s \n", film.genre)
		fmt.Printf("year : %d \n", film.year)
		fmt.Println()
	}
}

func main() {
	dataFilm := []movie{}
	addMovie("LOTR", "action", 120, 1999, &dataFilm)
	addMovie("avenger", "action", 120, 2019, &dataFilm)
	addMovie("spiderman", "action", 120, 2004, &dataFilm)
	addMovie("juon", "horror", 120, 2004, &dataFilm)

	displayMovies(&dataFilm)
}
