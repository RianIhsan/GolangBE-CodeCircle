package main

import "fmt"

func addToBuah(slice *[]string, elements ...string) {
	*slice = append(*slice, elements...)
}

func main() {
	var buah = []string{}

	addToBuah(&buah, "Jeruk")
	addToBuah(&buah, "Semangka")
	addToBuah(&buah, "Mangga")
	addToBuah(&buah, "Strawberry")
	addToBuah(&buah, "Durian")
	addToBuah(&buah, "Manggis")
	addToBuah(&buah, "Alpukat")

	for i, item := range buah {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
