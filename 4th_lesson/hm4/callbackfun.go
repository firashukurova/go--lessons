package main

import "fmt"

func main() {
	multiply := CreateMultiplier(6)
	fmt.Println(multiply(2))

	greeter := CreateGreeter("Hello")
	fmt.Println(greeter("Fira"))

	validator := CreateValidator("password123")
	fmt.Println(validator("123"))
}

//Создайте функцию CreateMultiplier, которая принимает число и возвращает функцию, умножающую переданное число на модуль этого числа.
//Используйте условные операторы внутри возвращаемой функции для проверки значений чисел

func CreateMultiplier(n int) func(int) int {
	if n < 0 {
		n = -n
	}
	return func(x int) int {
		return x * n
	}
}

//Создайте функцию CreateGreeter, которая принимает строку и возвращает функцию, которая принимает имя и возвращает приветствие с использованием переданной строки.
//Если строка пустая, то функция должна возвращать приветствие гостя. Используйте условные операторы внутри возвращаемой функции для проверки значений строк

func CreateGreeter(greeting string) func(string) string {
	return func(name string) string {
		if greeting == "" {
			return "Привет, гость " + name + "!"
		}
		return greeting + ", " + name + "!"
	}
}

//Создайте функцию CreateValidator, которая принимает строку пароля и возвращает функцию, которая принимает строку и возвращает true, если строка совпадает с паролем,
//и false в противном случае. Используйте условные операторы внутри возвращаемой функции для проверки значений строк

func CreateValidator(password string) func(string) bool {
	return func(input string) bool {
		return input == password
	}
}
