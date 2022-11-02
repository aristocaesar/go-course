package main

import "fmt"

func main() {
	const name string = "aristo caesar pratama"
	fmt.Println(name)

	// multiple declaration
	const (
		date string  = "19/08/01"
		pi   float32 = 3.14
	)
	fmt.Println(date)
	fmt.Println(pi)
}
