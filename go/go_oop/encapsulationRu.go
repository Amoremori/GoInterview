package main

import (
	"fmt"
)

// Определение структуры
type Circle1 struct {
	radius float64
}

// Методы для структуры Circle
func (c Circle1) SetRadius(radius float64) {
	c.radius = radius
}

func (c Circle1) GetRadius() float64 {
	return c.radius
}

func (c Circle1) CalculateArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// Создание экземпляра структуры Circle
	c := Circle1{radius: 5.0}

	// Изменение значения радиуса с помощью метода SetRadius
	c.SetRadius(7.0)

	// Получение значения радиуса с помощью метода GetRadius
	fmt.Println("Радиус окружности:", c.GetRadius())

	// Вычисление площади окружности с помощью метода CalculateArea
	fmt.Println("Площадь окружности:", c.CalculateArea())
}
