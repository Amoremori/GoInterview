package main

import "fmt"

// Shape Определяем интерфейс для геометрических фигур
type Shape interface {
	Area() float64
}

// Rectangle Определяем структуру для прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

// Реализуем метод Area() для структуры Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Определяем структуру для круга
type Circle struct {
	Radius float64
}

// Реализуем метод Area() для структуры Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// Создаем экземпляр прямоугольника
	rectangle := Rectangle{Width: 5, Height: 10}
	// Создаем экземпляр круга
	circle := Circle{Radius: 3}

	// Используем полиморфизм: вызываем метод Area() для различных фигур через интерфейс Shape
	fmt.Println("Площадь прямоугольника:", getArea(rectangle))
	fmt.Println("Площадь круга:", getArea(circle))
}

// Функция, использующая интерфейс Shape для вызова метода Area() на различных фигурах
func getArea(shape Shape) float64 {
	return shape.Area()
}
