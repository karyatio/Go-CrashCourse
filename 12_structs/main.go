package main

import (
	"fmt"
	"strconv"
)

// Person Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Person Greeting method(value reciever)
func (person Person) greet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName + " and i am " + strconv.Itoa(person.age)
}

// Person hasBirthday method(pointer reciever)
func (person *Person) hasBirthday() {
	person.age++
}

// Person getMarried method(pointer reciever)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "M" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Tio", lastName: "Saputra", city: "Bandar Lampung", gender: "M", age: 22}
	person2 := Person{"Elfin", "Sanjaya", "Bandar Lampung", "M", 23}
	person3 := Person{"Erika", "Sah", "Jakarta", "F", 20}

	// fmt.Println(person1)
	// fmt.Println(person2.firstName)

	// person2.age++
	// fmt.Println(person2)
	person1.hasBirthday()
	person3.getMarried(person2.lastName)
	person1.getMarried("LOL")

	fmt.Println(person1)
	fmt.Println(person2.greet())
	fmt.Println(person3)
}
