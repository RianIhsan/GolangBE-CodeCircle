package main

import "fmt"

func main() {
	sayuran := []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}

	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}
}
