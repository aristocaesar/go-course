package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Aristo Caesar Pratama",
		"addreas": "Banyuwangi",
	}

	fmt.Println(person)
	fmt.Println(len(person))
	delete(person, "addreas")
	person["age"] = "20"
	fmt.Println(person)
}
