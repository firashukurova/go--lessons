package main

import "fmt"

//3 Описание: Реализуйте структуру Circle с полем radius. Реализуйте методы Area и Circumference для вычисления площади и периметра круга.
//Методы:
//Area() float64
//Circumference() float64

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	pi := 3.14
	return pi * c.radius * c.radius
}

func (c *Circle) Circumference() float64 {
	pi := 3.14
	return 2 * pi * c.radius

}

func main() {

	circle := Circle{radius: 5}
	fmt.Printf("Площадь круга: %.2f\n", circle.Area())
	fmt.Printf("Длина окружности: %.2f\n", circle.Circumference())
}
