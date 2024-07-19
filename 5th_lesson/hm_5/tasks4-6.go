package main

import "fmt"

func main() {
	Fibonachi()
	//Euclid()
	//FizzBuzz()
}

//Напишите программу, которая выводит первые 10 чисел последовательности Фибоначчи

func Fibonachi() {
	var a int = 0
	var b int = 1
	fmt.Print(a, " ", b, " ")
	for i := 3; i <= 10; i++ {
		sum := a + b
		a = b
		b = sum
		fmt.Print(sum, " ")
	}
	fmt.Println()
}

//Напишите программу, которая находит наибольший общий делитель (НОД) двух чисел, используя алгоритм Евклида

func Euclid() {
	var num1 int
	var num2 int
	fmt.Println("Введите числа для вычисления НОД")
	fmt.Scan(&num1, &num2)

	//цикл while в Go
	for num2 != 0 && num1 != 0 {
		num1, num2 = num2, num1%num2
	}

	fmt.Printf("НОД %d", num1)
}

//Напишите программу, которая выводит числа от 1 до 100, заменяя числа, кратные 3, на "Fizz", числа, кратные 5, на "Buzz", а числа, кратные 3 и 5, на "FizzBuzz"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
