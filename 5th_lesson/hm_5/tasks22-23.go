package main

import "fmt"

func main() {

	//PrintNumbersUntilSum200()
	PerfectNumber()

	//Напишите программу, которая бесконечно запрашивает у пользователя ввод числа и прекращает выполнение, если введенное число делится на 7
	for {
		var input int
		fmt.Print("Enter number: ")
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
	fmt.Printf("\nGet sum: %d\n", sum)
}

//Напишите программу, которая находит все совершенные числа от 1 до n. (Совершенное число - это число, равное сумме всех своих делителей, кроме самого себя)

func PerfectNumber() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)

	fmt.Printf("Совершенные числа от 1 до %d:\n", n)
	for i := 1; i <= n; i++ {
		if isPerfect(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func isPerfect(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}
