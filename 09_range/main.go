package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 43, 65, 34}

	// Loop through id
	for i, id := range ids {
		fmt.Printf("%d -ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Printf("Total id : %d\n", sum)

	// Range with map
	emails := map[string]string{"Bobs": "bob@gmail.com", "tio": "karya.tiosaputra@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
