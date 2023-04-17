package main

import "fmt"

// Define an interface for geometric shapes
type ShapeEn interface {
	AreaEn() float64
}

// Define a struct for a rectangle
type RectangleEn struct {
	Width  float64
	Height float64
}

// Implement the Area() method for the Rectangle struct
func (r RectangleEn) AreaEn() float64 {
	return r.Width * r.Height
}

// Define a struct for a circle
type CircleEn struct {
	Radius float64
}

// Implement the Area() method for the Circle struct
func (c CircleEn) AreaEn() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// Create an instance of a rectangle
	rectangle := RectangleEn{Width: 5, Height: 10}
	// Create an instance of a circle
	circle := CircleEn{Radius: 3}

	// Use polymorphism: call the Area() method on different shapes through the Shape interface
	fmt.Println("Rectangle area:", getAreaEn(rectangle))
	fmt.Println("Circle area:", getAreaEn(circle))
}

// A function that uses the Shape interface to call the Area() method on different shapes
func getAreaEn(shape ShapeEn) float64 {
	return shape.AreaEn()
}
