package main

import (
	"fmt"
	"strings"
)

type Car struct {
	Model string
	Year  int
	Price float32
}

type Book struct {
	Title  string
	Author string
	Year   int
}

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

// Напишите функцию, которая принимает срез Book и год, и возвращает новый срез с книгами, изданными после указанного года.

func FilterBooksByYear(books []Book, year int) []Book {
	var filteredBooks []Book
	for _, book := range books {
		if book.Year > year {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

// Напишите функцию, которая принимает срез Employee и возвращает сотрудника с наибольшим возрастом.

func FindOldestEmployee(employees []Employee) Employee {
	oldest := employees[0]
	for _, employee := range employees {
		if employee.Age > oldest.Age {
			oldest.Age = employee.Age
		}

	}
	return oldest
}

//Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, в котором каждое число возведено в квадрат.

func SquareArr(arr []int) (arr2 []int) {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = v * v
	}
	return result
}

//Реализуйте функцию, которая принимает срез строк и возвращает новый срез, содержащий строки в нижнем регистре.

func DownArr(str []string) []string {
	result := make([]string, len(str))
	for i, s := range str {
		result[i] = strings.ToLower(s)
	}
	return result
}

//Реализуйте функцию, которая принимает срез структур Car и возвращает структуру Car, у которой цена минимальна.

func ReturnCheapCar(cars []Car) Car {
	minPrice := cars[0]
	for _, c := range cars {
		if c.Price < minPrice.Price {
			minPrice = c
		}
	}
	return minPrice

}

func main() {
	fmt.Println(SquareArr([]int{1, 2, 3, 4, 5}))
	strings := []string{"Hello", "WORLD", "Golang"}
	fmt.Println(DownArr(strings))

	m := map[string]string{"uno": "one"}
	for x := range m {
		fmt.Println(x)
	}

	cars := []Car{
		{Model: "BMW", Year: 2000, Price: 25000},
		{Model: "Toyota", Year: 1999, Price: 22000},
		{Model: "Ford", Year: 2009, Price: 27000},
	}

	fmt.Println(ReturnCheapCar(cars))

	books := []Book{
		{Title: "Go Programming Language", Author: "Jane Doe", Year: 2010},
		{Title: "Go Programming Language", Author: "Jane Doe", Year: 2011},
		{Title: "Go Programming Language", Author: "Jane Doe", Year: 2012},
	}

	fmt.Println(FilterBooksByYear(books, 2011))

}
