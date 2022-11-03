package main

import "fmt"

func main() {

	ipk := -1

	if ipk >= 0 {
		fmt.Println("Kamu lulus dengan IPK,", ipk)
	} else {
		fmt.Println("Mohon maaf nilai IPK tidak cukup untuk lulus")
	}

	name := "Aristo Caesar"

	if length := len(name); length >= 6 {
		fmt.Println("Nama terlalu panjang")
	}
	if len(name) >= 6 {
		fmt.Println("Nama terlalu panjang")
	}

}
