package main

import "fmt"

func main() {

	//Pyramid()
	//MultiplicationTable()
	PascalsTriangle()

}

func Pyramid() {
	var height int
	fmt.Println("Enter a height: ")
	fmt.Scan(&height)

	for i := 1; i <= height; i++ {
		// Печать пробелов
		for j := i; j <= height-1; j++ {
			fmt.Print(" ")
		}

		// Печать звездочек
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		// Переход на новую строку
		fmt.Println()
	}
}

//Напишите программу, которая выводит таблицу умножения от 1 до n, где n - введенное пользователем число

func MultiplicationTable() {
	var n int
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("**************")
	}
}

//Напишите программу, которая выводит треугольник Паскаля высотой n, где n - введенное пользователем число

func PascalsTriangle() {
	var h int
	fmt.Print("Enter height of triangle: ")
	fmt.Scan(&h)

	for i := 0; i < h; i++ {
		number := 1
		// Печатаем пробелы для выравнивания
		for space := 1; space <= h-i; space++ {
			fmt.Print("  ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("%4d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
