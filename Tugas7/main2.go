package main

import (
	"errors"
	"fmt"
)

func kelilingSegitigaSamaSisi(sisi float64, showString bool) (string, error) {
	if sisi == 0 {
		if showString {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic:", r)
				}
			}()

			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}

	keliling := sisi * 3

	if showString {
		result := fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %.2f cm adalah %.2f cm", sisi, keliling)
		return result, nil
	} else {
		return fmt.Sprintf("%.2f", keliling), nil
	}
}

func main() {
	result1, err1 := kelilingSegitigaSamaSisi(4, true)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := kelilingSegitigaSamaSisi(8, false)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(result2)
	}

	result3, err3 := kelilingSegitigaSamaSisi(0, true)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println(result3)
	}

	result4, err4 := kelilingSegitigaSamaSisi(0, false)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Println(result4)
	}
}
