package main

import "fmt"

func main() {

	name := ""

	switch {
	case name == "":
		fmt.Println("Nama Kosong")
	case len(name) <= 3:
		fmt.Println("Karakter terlalu sedikit!")
	case len(name) >= 6:
		fmt.Println("Karakter terlalu banyak!")
	}

}
