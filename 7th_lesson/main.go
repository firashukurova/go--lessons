package main

import "fmt"

type Age int
type Number int
type Score int

type Count int
type Temperature float32
type Price float64

type Adder func(int, int) int
type Operation func(int, int) int
type Comparator func(int, int) bool
type UnaryOperation func(int) int

func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}

type Address struct {
	City  string
	State string
}

type Person2 struct {
	Name string
	Age  int
	Addr Address
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Coordinate struct {
	X, Y int
}

func main() {
	//var myAge Age = 29
	//fmt.Println(myAge)

	p := new(Person)
	p.name = "Firuza"
	p.age = 29
	fmt.Println(p.name)
	fmt.Println(p.name)
	fmt.Println(p.name, p.age)

	p1 := Person2{
		Name: "Firuza",
		Age:  29,
		Addr: Address{
			City:  "Berlin",
			State: "Germany",
		},
	}

	fmt.Println(p1)

	var myAge Age = 29
	result := CheckAge(myAge)
	fmt.Println(result)

	var myAge2 Age = 17
	result = CheckAge(myAge2)
	fmt.Println(result)

	var num Number
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	res := ParityCheck(num)
	fmt.Println(res)

	var score Score
	fmt.Println("Введите число")
	fmt.Scanln(&score)
	r := RangeCheck(score)
	fmt.Println(r)

	var startCount Count = 10
	Countdown(startCount)
}

// Определение возраста совершеннолетия Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.

func CheckAge(age Age) string {
	if age >= 18 {
		return fmt.Sprintf("Вы совершенолетний, Ваш возраст %d", age)
	} else {
		return fmt.Sprintf("Вы не совершенолетний, Ваш возраст %d", age)
	}

}

//Проверка на четность. Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том,
//является ли оно четным или нечетным

func ParityCheck(number Number) string {
	if number%2 == 0 {
		return fmt.Sprintf("Число %d является четным", number)
	} else {
		return fmt.Sprintf("Число %d не является четным", number)
	}
}

//Проверка допустимого диапазона. Определите тип Score на основе int. Напишите функцию, которая принимает оценку и возвращает сообщение,
//находится ли она в допустимом диапазоне от 0 до 100

func RangeCheck(score Score) string {
	if score > 0 && score <= 100 {
		return fmt.Sprintf("Оценка %d в допустимом диапазоне", score)
	} else {
		return "Не допустимый диапазон"
	}
}

//Арифметические операции. Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения,
//вычитания и умножения. Используйте тип Operation для вызова этих функций

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}

func Division(a, b int) int {
	return a / b
}

// Сравнение чисел. Определите тип функции Comparator, которая принимает два int и возвращает bool. Напишите функции для проверки
// равенства и сравнения больше/меньше. Используйте тип Comparator для вызова этих функций.

func Comparison(num1, num2 int) bool {
	return num1 == num2
}

// Применение функции к числу. Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функции для удвоения
// и утроения числа. Используйте тип UnaryOperation для вызова этих функций.

func UnOperation2(num2 int) int {
	return num2 * 2
}

func UnOperation3(num3 int) int {
	return num3 * 3
}

// Обратный отсчет. Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля

func Countdown(count Count) {
	for i := int(count); i >= 0; i-- {
		fmt.Print(i)
	}
}

// Проверка температуры. Создайте псевдоним для типа float64 под названием Temperature. Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).

func CheckTemperature(temperature Temperature) string {
	if temperature > 0 {
		return "больше O"
	} else if temperature < 0 {
		return "меньше"
	} else {
		return "0"
	}
}

// Применение скидки. Создайте псевдоним для типа float64 под названием Price. Напишите функцию, которая принимает Price и возвращает новую цену с 10% скидкой.

func Discount(price Price) Price {
	discountAmount := price * 0.10
	discountPrice := price - discountAmount
	return Price(discountPrice)
}

// Информация о книге. Создайте структуру Book с полями Title (строка), Author (строка) и Pages (целое число). Напишите функцию, которая принимает Book и выводит информацию о книге.

func GetAuthor(book Book) {
	fmt.Printf("Название: %s\n", book.Title)
	fmt.Printf("Автор: %s\n", book.Author)
	fmt.Printf("Количество страниц: %d\n", book.Pages)

}

// Сравнение двух структур через указатели. Создайте структуру Coordinate с полями X и Y. Напишите функцию, которая принимает указатели на две
// Coordinate и возвращает сообщение о том, равны ли они или нет.

func CompareCoordinates(c1 *Coordinate, c2 *Coordinate) string {
	if c1.X == c2.X && c1.Y == c2.Y {
		return "Координаты равны"
	}
	return "Координаты не равны"
}
