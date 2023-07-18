package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		switch {
		case i%2 == 1 && i%3 == 0:
			fmt.Printf("%d - I LOVE CODING\n", i)
		case i%2 == 1:
			fmt.Printf("%d - CODE\n", i)
		case i%2 == 0:
			fmt.Printf("%d - CIRCLE\n", i)
		}
	}
}
