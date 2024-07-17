package main

import "fmt"

func main() {
	fmt.Println(Add(-3, 5))
	fmt.Println(CompareStrings("window", "window"))
	fmt.Println(Max(9, 19))
}

//Создайте функцию Add, которая принимает два целых числа и возвращает сумму их модулей. Используйте условные операторы для проверки значений чисел

func Add(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

//Создайте функцию CompareStrings, которая принимает две строки и возвращает "равны", если строки одинаковы, и "не равны"
//в противном случае. Используйте условные операторы для сравнения строк

func CompareStrings(first, second string) bool {
	if first == second {
		return true
	} else {
		return false
	}
}

//Создайте функцию Max, которая принимает два целых числа и возвращает большее из них. Используйте условные операторы для сравнения чисел

func Max(number1, number2 int) int {
	if number1 > number2 {
		return number1
	} else {
		return number2
	}
}
