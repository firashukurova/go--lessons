package main

import (
	"fmt"
	"strings"
)

func main() {
	a := ApplyOperation(-10, 3, operation)
	fmt.Println(a)

	CheckCondition(20, condition)

	FormatAndPrint("Hi", formatFunc)
}

//Создайте функцию ApplyOperation, которая принимает два целых числа и функцию, которая выполняет операцию над модулями этих чисел.
//Примените переданную функцию и выведите результат. Используйте условные операторы внутри функции для проверки значений чисел

func ApplyOperation(n1, n2 int, operation func(int, int) int) int {
	if n1 < 0 {
		n1 = -n1
	}
	if n2 < 0 {
		n2 = -n2
	}
	return operation(n1, n2)
}

func operation(num1 int, num2 int) int {
	return num1 + num2
}

//Создайте функцию CheckCondition, которая принимает целое число и функцию, возвращающую true или false.
//Если функция возвращает true, выведите "Условие выполнено", иначе "Условие не выполнено". Используйте условные операторы для проверки значений

func CheckCondition(q int, condition func(int) bool) {
	if condition(q) {
		fmt.Println("Условие выполнено")
	} else {
		fmt.Println("Условие не выполнено")
	}
}

func condition(n int) bool {
	if n > 0 {
		return true
	} else {
		return false
	}

}

//Создайте функцию FormatAndPrint, которая принимает строку и функцию, возвращающую отформатированную строку. Если строка пустая,
//callback – функция должна вывести “Пустая строка”, в противном случае, вывести строку которая получается от предыдущей преобразованием в заглавную всех символов.
//Примените переданную функцию и выведите результат. Используйте условные операторы для проверки значений строк

func FormatAndPrint(s string, formatFunc func(string) string) {
	if s == "" {
		fmt.Println("Пустая строка")
	} else {
		formattedString := formatFunc(s)
		fmt.Println(formattedString)
	}
}

func formatFunc(s string) string {
	s = "qwerty"
	return strings.ToUpper(s)
}
