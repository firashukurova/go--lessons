package main

import (
	"fmt"
	"strings"
)

func main() {

	//Создайте переменную operation, которая является функцией, принимающей два целых числа и возвращающей их разность.
	//Используйте условные операторы внутри функции для проверки значений чисел

	operation := func(num1, num2 int) int {
		if num1 < num2 {
			return num2 - num1
		} else {
			return num1 - num2
		}
	}
	fmt.Println(operation(23, 19))

	//Создайте переменную concat, которая является функцией, принимающей две строки и возвращающей их конкатенацию с учетом того,
	//что если строки не пустые добавить между ними разделение по пробелу. Используйте условные операторы внутри функции для проверки значений строк

	concat := func(str1, str2 string) string {
		if str1 != "" && str2 != "" {
			return str1 + " " + str2
		} else {
			return str1 + str2
		}
	}
	fmt.Println(concat("hello", "world"))

	//Создайте переменную multiply, которая является функцией, принимающей два целых числа и возвращающей произведение их модулей.
	//Используйте условные операторы внутри функции для проверки значений чисел

	multiply := func(nm1, nm2 int) int {
		if nm1 < 0 {
			nm1 = -nm1
		}
		if nm2 < 0 {
			nm2 = -nm2
		}
		return nm1 * nm2
	}
	fmt.Println(multiply(3, -4))

	CheckCondition(5, func(n int) bool {
		return n > 0
	})

	FormatAndPrint("hello", func(s string) string {
		return strings.ToUpper(s)
	})

}
