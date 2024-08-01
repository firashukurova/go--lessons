package main

import "fmt"

type Rectangle struct {
	width, height float64
}
type Counter struct {
	value int
}

type Logger struct{}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Describe() {
	fmt.Printf("Width: %v, Height: %v, Area: %v, Perimeter: %v\n", r.width, r.height, r.Area(), r.Perimeter())
}
func (Logger) Log(messages ...string) {
	for _, message := range messages {
		fmt.Println("Log:", message)
	}
}

func (c *Counter) Incrementer() func() int {
	return func() int {
		c.value++
		return c.value
	}
}

func main() {
	r := Rectangle{width: 10, height: 5}
	fmt.Println(r.Area())
	r.Scale(2)
	fmt.Println(r.Area())

	rect := Rectangle{width: 10, height: 5}
	rect.Describe() // Output: Width: 10, Height: 5, Area: 50, Perimeter: 30

	counter := Counter{}
	inc := counter.Incrementer()
	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3

	//Чем отличается рекурсия от замыкания?
}
