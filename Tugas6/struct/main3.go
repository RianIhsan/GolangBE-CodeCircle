package main

import "fmt"

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func addColor(p *phone, color string) {
	p.colors = append(p.colors, color)
}

func main() {
	myPhone := phone{
		name:  "iphone x",
		brand: "apple",
		year:  2018,
		colors: []string{
			"red",
			"blue,",
			"black,",
		},
	}

	fmt.Println("name :", myPhone.name)
	fmt.Println("brand :", myPhone.brand)
	fmt.Println("year :", myPhone.year)
	fmt.Println("colors :", myPhone.colors)

	addColor(&myPhone, "silver")

	fmt.Println("colors after adding silver :", myPhone.colors)
}
