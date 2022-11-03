package main

import "fmt"

func main() {

	// slice array merupakan sebuah cara untuk megambil index pada array
	months := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
		"sadtember",
	}

	fmt.Println(months[11:])

}
