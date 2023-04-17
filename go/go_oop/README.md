<h1 align="center">GoOOP(Ru-EN)</h1>

<span style="white-space: pre-line">1Ru.[Struct-Ru](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/structRu.go)</span>

<span style="white-space: pre-line">1En.[Struct-En](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/structEn.go)</span>

<span style="white-space: pre-line">2Ru.[Encapsulation-Ru](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/encapsulationRu.go)</span>

<h3>Encapsulation-Ru</h3>

```bash
package main

import (
"fmt"
)

// Определение структуры
type Circle struct {
radius float64
}

// Методы для структуры Circle
func (c Circle) SetRadius(radius float64) {
c.radius = radius
}

func (c Circle) GetRadius() float64 {
return c.radius
}

func (c Circle) CalculateArea() float64 {
return 3.14 * c.radius * c.radius
}

func main() {
// Создание экземпляра структуры Circle
c := Circle{radius: 5.0}

// Изменение значения радиуса с помощью метода SetRadius
c.SetRadius(7.0)

// Получение значения радиуса с помощью метода GetRadius
fmt.Println("Радиус окружности:", c.GetRadius())

// Вычисление площади окружности с помощью метода CalculateArea
fmt.Println("Площадь окружности:", c.CalculateArea())
}

```

<a>В этом примере структура Circle инкапсулирует данные о радиусе окружности и методы для работы с этими данными. Методы SetRadius, GetRadius и CalculateArea определены для структуры Circle и позволяют устанавливать значение радиуса, получать значение радиуса и вычислять площадь окружности соответственно. Доступ к данным о радиусе осуществляется только через методы, что обеспечивает инкапсуляцию данных и скрывает их от прямого доступа извне.</a>

<span style="white-space: pre-line">2En.[Encapsulation-En](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/encapsulationRu.go)</span>

<h3>Encapsulation-En</h3>

```bash
package main

import (
"fmt"
)

// Definition of the Circle structure
type Circle struct {
radius float64
}

// Methods for the Circle structure
func (c Circle) SetRadius(radius float64) {
c.radius = radius
}

func (c Circle) GetRadius() float64 {
return c.radius
}

func (c Circle) CalculateArea() float64 {
return 3.14 * c.radius * c.radius
}

func main() {
// Creating an instance of the Circle structure
c := Circle{radius: 5.0}

// Changing the radius value using the SetRadius method
c.SetRadius(7.0)

// Getting the radius value using the GetRadius method
fmt.Println("Circle radius:", c.GetRadius())

// Calculating the area of the circle using the CalculateArea method
fmt.Println("Circle area:", c.CalculateArea())
}


```

<a>In this example, the Circle structure encapsulates the data about the radius of a circle and the methods for working with this data. The SetRadius, GetRadius, and CalculateArea methods are defined for the Circle structure and allow setting the radius value, getting the radius value, and calculating the area of the circle, respectively. Access to the radius data is only possible through the methods, which ensures encapsulation of the data and hides it from direct access from outside.</a>

<span style="white-space: pre-line">3Ru.[Inheritance/Composition/Embedding-Ru](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/compositionRu.go)</span>

<h3>Inheritance/Composition/Embedding-Ru</h3>

<a>В Go основным механизмом для повторного использования кода является композиция, где вы можете встроить одну структуру внутрь другой, чтобы достичь формы "наследования". Этот подход также называется "эмбеддинг" в Go. Рассмотрим пример, чтобы проиллюстрировать этот концепт: </a>

```bash
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

```
<a>В приведенном выше примере у нас есть структура Animal с полями name и weight, и структура Dog, встраивающая структуру Animal. Это означает, что структура Dog автоматически наследует поля и методы структуры Animal. Вы можете обращаться к этим полям и методам напрямую из структуры Dog, как показано в функции main.</a>

<a>Строго говоря, Go предпочитает подход на основе композиции перед традиционным наследованием, чтобы сохранить простоту языка и его легкость в понимании. Встраивание позволяет повторно использовать код без некоторых сложностей и проблем, связанных с наследованием в других языках.</a>

<span style="white-space: pre-line">3En.[Inheritance/Composition/Embedding-EN](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/compositionEn.go)</span>

<span style="white-space: pre-line">4Ru.[Polymorphism-Ru](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/polymorphismRu.go)</span>

<h3>PolymorphismRu</h3>


```bash
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

```

<a>В этом примере определен интерфейс Shape, который имеет метод Area(). Затем две структуры, Rectangle и Circle, реализуют этот интерфейс, предоставляя свои реализации метода Area(). Функция getArea() принимает аргумент типа Shape, и внутри нее вызывается метод Area() на переданной фигуре, будь то прямоугольник или круг, благодаря полиморфизму. Это позволяет использовать одну функцию для расчета площади различных геометрических фигур без явного указания их типа.</a>

<span style="white-space: pre-line">4En.[Polymorphism-En](https://github.com/Amoremori/GoInterview/blob/main/go/go_oop/polymorphismEn.go)</span>

<h3>PolymorphismEn</h3>




```bash
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

```

<a>In this example, an interface called Shape is defined with a method Area(). Then two structs, Rectangle and Circle, implement this interface by providing their own implementations of the Area() method. The getArea() function takes an argument of type Shape, and inside it, the Area() method is called on the passed shape, whether it's a rectangle or a circle, thanks to polymorphism. This allows using a single function to calculate the area of different geometric shapes without explicitly specifying their type.</a>