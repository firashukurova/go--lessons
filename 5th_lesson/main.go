package main

import "fmt"

func main() {
	//for i := 0; i <= 9; i++ {
	//	fmt.Println(i)
	//}

	//a := 0
	//for ; a < 10; a++ {
	//
	//}
	//var b int
	//for b = 0; b <= 9; b++ {
	//	fmt.Println(b)
	//}
	//
	//fmt.Println(b)
	//
	//var num = 546
	//for num > 0 {
	//	a := num % 10
	//	fmt.Println(a)
	//	num /= 10
	//}
	//
	//for n := 546; n > 0 && n > 10; n /=10{
	//	a := n % 10
	//	fmt.Println(a)
	//}

	//for i := 1; i < 100; i++ {
	//	if i%3 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	//for i := 1; i < 10; i++ {
	//	for j := 1; j < 10; j++ {
	//		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	//	}
	//	fmt.Println("-------------------------")
	//}

	//Дано целое число N (> 0). Найти сумму  1 + 1/2 + 1/3 + . . . + 1/N
	//var num = 5
	//var output float64 = 0
	//for i := float64(num); i <= float64(num); i++ {
	//	output += 1 / i
	//}
	//fmt.Println(output)
	//
	////Дано целое число N (> 0). Найти сумму  N2 + (N + 1)2 + (N + 2)2 + . . . + (2·N)2  (целое число)
	//
	//var N = 6
	//var sum int
	//for i := N; i <= 2*N; i++ {
	//	sum = sum + i*i
	//}
	//fmt.Println(sum)

	// Дано целое число N (> 0). Найти произведение  1.1 · 1.2 · 1.3 · . . .  (N сомножителей)
	//Даны десять вещественных чисел. Найти их произведение
	//
	//var num1 float64 = 1
	//var num2 float64
	//for i := 0; i < 10; i++ {
	//	fmt.Scanln(&num2)
	//	num1 = num1 * num2
	//}
	//fmt.Println(num1)

	//Дано целое число N и набор из N положительных вещественных чисел. Вывести в том же порядке дробные части всех чисел из данного набора
	//(как вещественные числа с нулевой целой частью), а также произведение всех дробных частей

	var N int
	var num, fractionalPart, product float64

	fmt.Print("Введите количество чисел N: ")
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&num)
		fractionalPart = num - float64(int(num))
		fmt.Printf("Дробная часть: %.2f\n", fractionalPart)
		product += fractionalPart
	}

	fmt.Printf("Произведение дробных частей: %.2f\n", product)
}
