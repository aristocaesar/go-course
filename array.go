package main

import "fmt"

func main() {
	var students [3]string

	students[0] = "Aristo Caesar Pratama"
	students[1] = "Hana Wulan Agusta"
	students[2] = "Nova Kharisama"

	var cars = [3]string{
		"Pajero",
		"Civic Turbo",
		"Alpard",
	}

	fmt.Println(students)
	fmt.Println(cars)
	fmt.Println(cars[2])
}
