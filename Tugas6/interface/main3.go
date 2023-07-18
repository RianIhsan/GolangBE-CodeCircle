package main

import (
	"fmt"
)

func luasPersegi(sisi int, tampilkanKeterangan bool) interface{} {
	if sisi == 0 {
		if tampilkanKeterangan {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}

	luas := sisi * sisi

	if tampilkanKeterangan {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return luas
	}
}

func main() {
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
}
