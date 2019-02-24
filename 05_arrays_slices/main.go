package main

import "fmt"

func main() {
	// // Array
	// var fruitArray [2]string

	// // Assign value
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)

	// SLICES
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
