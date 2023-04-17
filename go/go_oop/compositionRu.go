package main

import "fmt"

// Структура Animal
type Animal1 struct {
	name   string
	weight float64
}

// Структура Dog, встраивающая структуру Animal
type Dog1 struct {
	Animal1
	breed string
}

// Метод для Animal
func (a Animal1) Eat1() {
	fmt.Println("Животное ест.")
}

// Метод для Dog
func (d Dog1) Bark1() {
	fmt.Println("Собака лает.")
}

func main() {
	// Создаем экземпляр структуры Dog
	dog := Dog1{
		Animal1: Animal1{
			name:   "Бадди",
			weight: 25.5,
		},
		breed: "Лабрадор",
	}

	// Обращаемся к полям и методам как из структуры Animal, так и из структуры Dog
	fmt.Println("Имя собаки:", dog.name)
	fmt.Println("Вес собаки:", dog.weight)
	dog.Eat1()
	dog.Bark1()
}
