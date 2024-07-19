package main

import "fmt"

func main() {

	//PositiveNumber()
	//ProductOfNumbers()
	Palindrome()
}

//Напишите программу, которая запрашивает у пользователя ввод положительного числа и повторяет запрос, пока не будет введено положительное число

func PositiveNumber() {
	var number int
	for {
		fmt.Println("Add Positive Number")
		fmt.Scan(&number)
		if number > 0 {
			fmt.Println("Positive Number")
			return
		} else {
			fmt.Println("No Positive Number")
		}
	}
}

//Напишите программу, которая вычисляет произведение всех чисел от 1 до введённого числа n, но прекращает выполнение, если результат превышает 1000

func ProductOfNumbers() {
	var num int
	var total = 1
	fmt.Println("Add Number")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		if total > 1000 {
			break
		}
		total *= i
	}
	if total > 1000 {
		fmt.Println("Out of range")
	} else {
		fmt.Printf("Result: %d\n", total)
	}
}

//Напишите программу, которая запрашивает у пользователя ввод числа и проверяет, является ли оно палиндромом

func Palindrome() {
	var input int
	fmt.Println("Add Number")
	fmt.Scan(&input)

	original := input
	reversed := 0

	for input > 0 {
		digit := input % 10
		reversed = reversed*10 + digit
		input /= 10
	}
	if original == reversed {
		fmt.Println(original, "is palindrome")
	} else {
		fmt.Println(original, "isn't palindrome")
	}
}
