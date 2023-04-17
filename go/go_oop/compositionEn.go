package main

import "fmt"

// Animal struct
type Animal struct {
	name   string
	weight float64
}

// Dog struct which embeds Animal
type Dog struct {
	Animal
	breed string
}

// Method for Animal
func (a Animal) Eat() {
	fmt.Println("The animal is eating.")
}

// Method for Dog
func (d Dog) Bark() {
	fmt.Println("The dog is barking.")
}

func main() {
	// Create a Dog instance
	dog := Dog{
		Animal: Animal{
			name:   "Buddy",
			weight: 25.5,
		},
		breed: "Labrador",
	}

	// Access fields and methods from both Animal and Dog
	fmt.Println("Dog Name:", dog.name)
	fmt.Println("Dog Weight:", dog.weight)
	dog.Eat()
	dog.Bark()
}
