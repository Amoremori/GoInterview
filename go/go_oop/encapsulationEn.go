package main

import (
	"fmt"
)

// Definition of the Circle structure
type Circle2 struct {
	radius float64
}

// Methods for the Circle structure
func (c Circle2) SetRadius(radius float64) {
	c.radius = radius
}

func (c Circle2) GetRadius() float64 {
	return c.radius
}

func (c Circle2) CalculateArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// Creating an instance of the Circle structure
	c := Circle2{radius: 5.0}

	// Changing the radius value using the SetRadius method
	c.SetRadius(7.0)

	// Getting the radius value using the GetRadius method
	fmt.Println("Circle radius:", c.GetRadius())

	// Calculating the area of the circle using the CalculateArea method
	fmt.Println("Circle area:", c.CalculateArea())
}
