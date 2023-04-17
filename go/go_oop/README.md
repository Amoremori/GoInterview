<h1 align="center">GoOOP</h1>

<span style="white-space: pre-line">1.[Struct](https://github.com/Amoremori/GoInterview/blob/main/go/go_types/interfaces.go)</span>

<span style="white-space: pre-line">2.[Encapsulation](https://github.com/Amoremori/GoInterview/blob/main/go/go_types/interfaces.go)</span>

<span style="white-space: pre-line">3.[Inheritance/Composition/Embedding](https://github.com/Amoremori/GoInterview/blob/main/go/go_types/interfaces.go)</span>

<span style="white-space: pre-line">4.[Polymorphism](https://github.com/Amoremori/GoInterview/blob/main/go/go_types/interfaces.go)</span>
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