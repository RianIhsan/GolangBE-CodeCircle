package main

import "fmt"

func main() {
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	sliceKalimat := kalimat[2:7]

	for _, kata := range sliceKalimat {
		fmt.Printf("%s ", kata)
	}
}
