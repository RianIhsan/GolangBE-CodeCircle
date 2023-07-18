package main

import "fmt"

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type dataHp interface {
	android()
	apple()
}

func (sPhone *phone) android() string {
	dataAndroid := fmt.Sprintf("name : %s \nBrand : %s \nYear : %d \nColors : %v\n", sPhone.name, sPhone.brand, sPhone.year, sPhone.colors)
	return dataAndroid
}

func (aPhone *phone) apple() string {
	dataApple := fmt.Sprintf("name : %s \nBrand : %s \nYear : %d \nColors : %v\n", aPhone.name, aPhone.brand, aPhone.year, aPhone.colors)
	return dataApple
}
func main() {
	sAndroid := phone{
		name:   "J5",
		brand:  "samsung",
		year:   2021,
		colors: []string{"black", "red", "blue"},
	}

	fmt.Println(sAndroid.android())

	sApple := phone{
		name:   "iphone x",
		brand:  "apple",
		year:   2021,
		colors: []string{"white", "pink", "yellow"},
	}

	fmt.Println(sApple.apple())
}
