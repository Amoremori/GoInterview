package main

import "fmt"

// Define a Person structure
type Person struct {
	name string
	age  int
}

// Method to get information about a person
func (p Person) getInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

func main() {
	// Create a new Person object
	p := Person{"John", 25}

	// Print information about the person using the getInfo() method
	fmt.Println(p.getInfo()) // Name: John, Age: 25
}
