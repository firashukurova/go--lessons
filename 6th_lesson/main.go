package main

import "fmt"

func main() {
	//var number int
	//fmt.Println("Enter a number: ")
	//fmt.Scan(&number)
	//if number%2 != 0 && number > 0 {
	//	fmt.Println("Positive Odd number")
	//}
	//if number < 0 {
	//	fmt.Println("Negative number")
	//} else if number == 0 {
	//	fmt.Println("Zero")
	//}

	//var x int = 10
	//var ptr *int
	//fmt.Println(ptr)
	//ptr = &x
	//fmt.Println("Значение х:", x)
	//fmt.Println("Адрес x:", &x)
	//fmt.Println("Значение ptr:", ptr)
	//fmt.Println("ptr указывает на переменную х:", *ptr) //разыменования

	//var m int = 10
	//fmt.Println("before", m)
	//UpdateValue(&m)
	//fmt.Println("after", m)

	//var a int = 5
	//var b int = 6
	//fmt.Println("before", a, b)
	//Swap(&a, &b)
	//fmt.Println("after", a, b)

	var ptr *int
	fmt.Println(ptr)

	//PrintValue(ptr)
	//fmt.Println("ptr", ptr)

	//var m = new(int)
}

func UpdateValue(num *int) {
	*num = *num + 10
}

//Напишите функцию swap, которая меняет местами значения двух переменных с использованием указателей

func Swap(a, b *int) {
	*a, *b = *b, *a
}

//Напишите функцию printValue, которая принимает указатель на целое число и выводит его значение. Если указатель равен nil,
//выведите сообщение "Указатель пуст"

func PrintValue(ptr *int) {
	if ptr == nil {
		fmt.Println("Указатель пуст")
	} else {
		fmt.Println("Значение", ptr)
	}

}

//Напишите программу для управления персональными данными. Реализуйте функции для хранения и вывода имени и возраста человека с использованием указателей

func SaveName(namePtr *string, name string) {
	*namePtr = name
}

func SaveAge(AgePtr *int, age int) {
	*AgePtr = age
}

//Напишите программу для учета голосов на выборах. Реализуйте функции для подсчета голосов за каждого кандидата и определения победителя

func Vote(candidatePtr *string, votesPtr *int) {
	*votesPtr++
}

func Winner(candidate1Votes *int, candidate2Votes *int) string {
	if *candidate1Votes > *candidate2Votes {
		return "Кандидат 1"
	} else if *candidate1Votes < *candidate2Votes {
		return "Кандидат 2"
	} else {
		return "Ничья"
	}
}

//Напишите программу для учета расходов на семейный бюджет. Реализуйте функции для добавления новой записи расходов и вывода общей суммы расходов

func AddExpense(total *float64, amount float64) {
	*total += amount
}

func PrintTotalExpense(totalPtr *float64) {
	fmt.Printf("Общаяя сумма расходов %.2f\n", *totalPtr)
}

// Напишите программу для конвертации суммы из долларов в евро по текущему курсу. Начальный курс равен 0.85 евро за доллар.
// Реализуйте функции для обновления курса валюты и конвертации суммы. Если конвертированная сумма превышает 5000 евро,
// выведите сообщение "Превышен лимит конвертации"
