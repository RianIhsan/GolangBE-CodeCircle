package main

import (
	"fmt"
)

func main() {
	var kelahiran = 2003
	if kelahiran > 1944 && kelahiran < 1964 {
		fmt.Println("Baby boomer, kelahiran 1944 s.d 1964")
	} else if kelahiran > 1965 && kelahiran < 1979 {
		fmt.Println("Generasi X. Kelahiran 1965 s.d 1979")
	} else if kelahiran > 1980 && kelahiran < 1994 {
		fmt.Println("Generasi Y (Millenials), kelahiran 1980 s.d 1994")
	} else if kelahiran > 1995 && kelahiran < 2015 {
		fmt.Println("Generasi Z, kelahiran 1995 s.d 2015")
	} else {
		fmt.Println("Input Tidak Valid")
	}
}
