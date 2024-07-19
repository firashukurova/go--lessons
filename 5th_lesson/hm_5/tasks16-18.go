package main

import "fmt"

func main() {

	Factorial()
	Sum()
	PrintNumbers()
}

//Напишите программу, которая бесконечно запрашивает у пользователя ввод числа и выводит его факториал, пока не будет введено число меньше 0

func Factorial() {
	var number int
	for {
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		if number < 0 {
			fmt.Println("Incorrect number")
			break
		}
		result := 1
		for i := 1; i <= number; i++ {
			result *= i
		}
		fmt.Printf("Factorial of %d is %d\n", number, result)
	}

}

//Напишите программу, которая бесконечно запрашивает у пользователя ввод двух чисел и выводит их сумму.

func Sum() {
	var a, b int
	for {
		fmt.Println("Enter a number: ")
		fmt.Scan(&a)
		fmt.Println("Enter b number: ")
		fmt.Scan(&b)
		result := a + b
		fmt.Println(result)
	}
}

//Напишите программу, которая выводит все числа от 1 до 100, но пропускает числа, которые являются квадратами

func PrintNumbers() {
	for i := 1; i <= 100; i++ {
		isSquare := false
		for j := 1; j*j <= i; j++ {
			if j*j == i {
				isSquare = true
				break
			}
		}
		if !isSquare {
			fmt.Println(i)
		}
	}
	fmt.Println()
}
