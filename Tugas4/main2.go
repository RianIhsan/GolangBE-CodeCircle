package main

import "fmt"

func main() {
	for i := 0; i <= 7; i++ {
		line := ""
		for j := 0; j <= i; j++ {
			line += "#"
		}
		fmt.Println(line)
	}
}
