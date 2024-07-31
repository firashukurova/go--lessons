package main

import "fmt"

func main() {
	//PrintPrimeNumbers()
	//PrintNumbers2()
	CheckFun()

}

//Напишите программу, которая выводит только те числа от 1 до 50, которые являются простыми

func PrintPrimeNumbers() {
	for i := 2; i <= 50; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

//Напишите программу, которая выводит все числа от 1 до 30, кроме чисел, которые делятся на 2 или 3

func PrintNumbers2() {
	for i := 1; i <= 30; i++ {
		if i%2 != 0 && i%3 != 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
}

//Напишите программу, которая выводит числа от 1 до 50, но прекращает выполнение, если встречает число, которое является кубом

func CheckFun() {
	for i := 1; i <= 50; i++ {
		fmt.Printf("%d ", i)
		for j := 1; j*j*j <= i; j++ {
			if j*j*j == i {
				fmt.Println("\nCube", i)
				return
			}
		}
	}
	fmt.Println("\nCube isn't find")
}
