package main

import "fmt"

func main() {
	var integer32 int32 = 128
	var integer64 int64 = int64(integer32)
	// jika nilai overflow, maka akan loop value
	var integer8 int8 = int8(integer64)
	fmt.Println(integer32)
	fmt.Println(integer64)
	fmt.Println(integer8)

	// byte to sring
	name := "Aristo Caesar Pratama"
	e := name[0]
	fmt.Println("Byte ", e, " is ", string(e))
}
