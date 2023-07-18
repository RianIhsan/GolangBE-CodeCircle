package main

import (
	"fmt"
)

func main() {
	var resultPrefix string = "Hasil penjumlahan dari\n"
	var firstNumbers []int = []int{6, 8}
	var secondNumbers []int = []int{12, 14}

	sum := 0
	for _, num := range firstNumbers {
		sum += num
	}
	for _, num := range secondNumbers {
		sum += num
	}

	output := fmt.Sprintf("%s", resultPrefix)
	for i, num := range firstNumbers {
		output += fmt.Sprintf("%d", num)
		if i < len(firstNumbers)-1 || len(secondNumbers) > 0 {
			output += " + "
		}
	}
	for i, num := range secondNumbers {
		output += fmt.Sprintf("%d", num)
		if i < len(secondNumbers)-1 {
			output += " + "
		}
	}
	output += fmt.Sprintf(" = %d", sum)

	fmt.Println(output)
}
