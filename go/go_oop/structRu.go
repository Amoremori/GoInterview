package main

import "fmt"

// Определим структуру Person
type Person1 struct {
	name string
	age  int
}

// Метод для получения информации о персоне
func (p Person1) getInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

func main() {
	// Создадим новый объект Person
	p := Person1{"John", 25}

	// Выведем информацию о персоне, используя метод getInfo()
	fmt.Println(p.getInfo()) // Name: John, Age: 25
}
