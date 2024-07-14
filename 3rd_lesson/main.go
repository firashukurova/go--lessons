package main

import "fmt"

func main() {

	//a := 5
	//b := 6
	//c := a + b
	//fmt.Println(c)
	//
	//a = 7
	//b = 8
	//c = a + b
	//fmt.Println(c)
	//
	////fmt.Println(sum(4, 5))
	//// fmt.Println(sum(4.5, 5))
	//printHello("Fira")

	//e := 6
	//d := float64(e)
	//fmt.Println(d)

	f := sum1(8)
	i := f()
	fmt.Println(i)

	//sumAndMultiply(3,4)
	fmt.Println(sumAndMultiply(5, 6))
	//функция как переменная
	//безымянные функции анонимные функции
	//callback

}

func sum1(a int) func() int {
	b := 6
	return func() int {
		return a + b
	}
}

func sum(a, b int) int {
	c := a + b
	return c
}

func printHello(name string) {
	fmt.Println("Hello", name)
}

func inc(a int) {
	a++
	fmt.Println(a)
}

func sumAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}
