package main

import (
	"fmt"
	"strings"
)

func main() {
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
	fmt.Println(GetWelcomeMessage())
	fmt.Println(GetPiValue())
	IsServerActive()
	PrintNumber(32)
	GreetUser("Fira")
	ToggleLight(true)
	fmt.Println(Add(3, 4))
	fmt.Println(Concat("Cat", "Dog"))
	fmt.Println(IsEven(31))

	//Создайте переменную adder, которая является функцией, принимающей два целых числа и возвращающей их сумму.
	adder := func(number1, number2 int) int {
		return number1 + number2
	}
	fmt.Println(adder(10, 25))

	//Создайте переменную concatenator, которая является функцией, принимающей две строки и возвращающей их конкатенацию
	concatenator := func(str1, str2 string) string {
		return str1 + str2
	}
	fmt.Println(concatenator("Go", "Lang"))

	//Создайте переменную isPositive, которая является функцией, принимающей целое число и возвращающей true, если число положительное
	isPositive := func(somenum int) bool {
		return somenum > 0
	}
	fmt.Println(isPositive(-9))
	fmt.Println(Calculate(9, 3, Add))
	Execute(true, ServerStatus)

	fmt.Println(Apply(5, Fun))
	fmt.Println(StringRepeater(2)("Win"))

	truePrint := ConditionalPrinter(true)
	truePrint("This is true")
}

// Создайте функцию PrintGreeting, которая печатает "Hello, World!" на экран

func PrintGreeting() {
	fmt.Println("Hello, World!")

}

// Создайте функцию DisplaySeparator, которая печатает строку из 20 символов - для разделения текста
func DisplaySeparator() {
	separator := ""
	for i := 0; i < 20; i++ {
		separator += "-"
	}
	fmt.Println(separator)
}

// Создайте функцию NotifyStart, которая выводит на экран сообщение "Program started"
func NotifyStart() {
	fmt.Println("Program started")
}

// Создайте функцию GetWelcomeMessage, которая возвращает строку "Welcome!"
func GetWelcomeMessage() string {
	w := "Welcome"
	return w
	//return "Welcome"
}

// Создайте функцию GetPiValue, которая возвращает значение числа π с точностью до двух знаков после запятой (3.14)
func GetPiValue() float64 {
	pi := 3.14
	return pi
	//return 3.14
}

// Создайте функцию IsServerActive, которая возвращает булево значение true
func IsServerActive() bool {
	return true
}

// Создайте функцию PrintNumber, которая принимает целое число и выводит его на экран
func PrintNumber(number int32) {
	fmt.Println(number)
}

// Создайте функцию GreetUser, которая принимает строку с именем пользователя и выводит приветствие
func GreetUser(name string) {
	fmt.Printf("Hello %s\n", name)
}

// Создайте функцию ToggleLight, которая принимает булевое значение и выводит "Light on" или "Light off" в зависимости от значения аргумента

func ToggleLight(light bool) {
	if light {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

//Создайте функцию Add, которая принимает два целых числа и возвращает их сумму

func Add(a, b int) int {
	return a + b
}

// Создайте функцию Concat, которая принимает две строки и возвращает их конкатенацию
func Concat(one, two string) string {
	return one + two
}

// Создайте функцию IsEven, которая принимает целое число и возвращает true, если число четное, и false в противном случае
func IsEven(num int) bool {
	return num%2 == 0
}

// Создайте функцию Calculate, которая принимает два целых числа и функцию для их обработки.
// Примените переданную функцию к этим числам и верните результат
func Calculate(n1 int, n2 int, Add func(int, int) int) int {
	return Add(n1, n2)
}

// Создайте функцию Execute, которая принимает булевое значение и функцию, которая принимает булевое значение и ничего не возвращает.
// Выполните переданную функцию с переданным булевым значением
func Execute(l bool, ServerStatus func(bool)) {
	ServerStatus(l)
}

func ServerStatus(status bool) {
	fmt.Println("Server is active:", status)
}

// Создайте функцию Apply, которая принимает целое число и функцию, которая принимает целое число и возвращает целое число.
// Примените переданную функцию к переданному числу и верните результат
func Apply(smn int, Fun func(int) int) int {
	return Fun(smn)
}

func Fun(i int) int {
	return i * i * i
}

// Создайте функцию Multiplier, которая принимает целое число factor и возвращает функцию, умножающую переданное ей число на factor
func Multiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

// Создайте функцию StringRepeater, которая принимает целое число n и возвращает функцию, повторяющую переданную ей строку n раз.
func StringRepeater(n3 int) func(string) string {
	return func(s string) string {
		return strings.Repeat(s, n3)
	}
}

// Создайте функцию ConditionalPrinter, которая принимает булево значение condition и возвращает функцию, печатающую строку, если condition истинно
func ConditionalPrinter(bool2 bool) func(st string) {
	return func(st string) {
		if bool2 {
			fmt.Println(st)
		}

	}
}
