package main

import "fmt"

func main() {
	// Using var
	var age int32 = 22
	var isCool = true

	name, email := "Tio Saputra", "karya.tiosaputra@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool)
}
