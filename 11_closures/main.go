package main

import "fmt"

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := addr()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
