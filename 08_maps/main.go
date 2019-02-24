package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bobs"] = "bob@gmail.com"
	// emails["Tio"] = "karya.tiosaputra@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Bobs": "bob@gmail.com", "tio": "karya.tiosaputra@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Tio"])

	// Delete form map
	delete(emails, "Bobs")
	fmt.Println(emails)
}
