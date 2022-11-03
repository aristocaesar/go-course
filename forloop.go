package main

import "fmt"

func main() {

	// for index := 1; index <= 10; index++ {
	// 	fmt.Println(index)
	// }

	mounths := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	for i := 0; len(mounths) > i; i++ {
		fmt.Println(mounths[i])
	}

	for index, mounth := range mounths {
		fmt.Println(index, "->", mounth)
	}

	for _, mounth := range mounths {
		fmt.Println(mounth)
	}

}
