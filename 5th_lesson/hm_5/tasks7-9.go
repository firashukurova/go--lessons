package main

import (
	"fmt"
)

func main() {
	//PrimeNumber()
	//Divisor()
	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)
	result := sumOfDigits(num)
	fmt.Printf("Сумма цифр числа %d равна %d\n", num, result)
}

//Напишите программу, которая проверяет, является ли число простым

func PrimeNumber() {
	fmt.Println("Add number")
	var number int
	fmt.Scan(&number)
	if number <= 1 {
		fmt.Println("Zero and number 1 is not prime")
		return
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			fmt.Printf("Number %d isn't prime\n", number)
			return

		}
	}
	fmt.Printf("Number %d is prime\n", number)
}

//func PrimeNumber() {
//	fmt.Println("Add number")
//	var number int
//	fmt.Scan(&number)
//
//	if number <= 1 {
//		fmt.Printf("Number %d is not prime\n", number)
//		return
//	}
//
//	isPrime := true
//	for i := 2; i < number; i++ {
//		if number % i == 0 {
//			isPrime = false
//			break
//		}
//	}
//
//	if isPrime {
//		fmt.Printf("Number %d is prime\n", number)
//	} else {
//		fmt.Printf("Number %d is not prime\n", number)
//	}
//}

//Напишите программу, которая выводит все делители числа

func Divisor() {
	fmt.Println("Add number")
	var divNumber int
	fmt.Scan(&divNumber)
	for i := 1; i <= divNumber; i++ {
		if divNumber%i == 0 {
			fmt.Printf("Divisors of %d is %d\n", divNumber, i)
		}
	}

}

//Напишите программу, которая находит сумму цифр числа

func sumOfDigits(number int) int {
	sum := 0
	for number > 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}
	return sum
}
