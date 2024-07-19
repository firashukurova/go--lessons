package main

import "fmt"

func main() {

	PrintNumbersUntilSum200()

	//Напишите программу, которая бесконечно запрашивает у пользователя ввод числа и прекращает выполнение, если введенное число делится на 7
	for {
		var input int
		fmt.Println("Enter number: ")
		fmt.Scan(&input)
		if input%7 == 0 {
			break
		}
		fmt.Println(input)
	}
}

//Напишите программу, которая выводит числа от 1 до 100, но прекращает выполнение, если сумма всех выведенных чисел превышает 200

func PrintNumbersUntilSum200() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 200 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nСумма достигла: %d\n", sum)
}
