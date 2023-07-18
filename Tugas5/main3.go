package main

import "fmt"

func fruits(name string, jenisBuah ...string) {
	fmt.Print("Hello nama saya " + name)
	fmt.Print("Buah kesukaan saya ")
	for _, buah := range jenisBuah {
		fmt.Print(buah)
	}
}

func main() {
	fruits("Rian ", "Rambutan ", "Durian ", "Pear ")
	fmt.Print()
}
