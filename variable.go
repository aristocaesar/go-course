package main

import "fmt"

func main() {
	var name string
	var umur int8 = 127
	name = "Aristo caesar"
	fmt.Println(name)
	fmt.Println(umur)

	// Agar tidak menggunakan var
	alamat := "Banyuwangi"
	fmt.Println(alamat)

	// mutitple declaration
	var (
		firstname = "aristo"
		lastname  = "caesar"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}
