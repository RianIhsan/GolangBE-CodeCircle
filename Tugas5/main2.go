package main

import "fmt"

func introduce(name, gender, hobi, age string) (string, string, string, string) {
	introName := "Pak " + name
	introGender := "berjenis kelamin " + gender
	introHobi := "dan seorang " + hobi
	introAge := "yang berusia " + age

	return introName, introGender, introHobi, introAge
}

func main() {
	fmt.Println(introduce("Rian", "Laki Laki", "Web Developer", "20"))
}
