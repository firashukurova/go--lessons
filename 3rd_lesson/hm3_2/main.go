package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//GetNumber(4, Add)
	//
	//result := ProcessNumber(5,
	//	func(x int) int { return x * 2 },
	//	func(x int) int { return x + 1 })
	//fmt.Println(result)

	password := Randomizer()
	fmt.Printf("Password: %s\n", password)

	code := GenerateInvoiceNumber()
	fmt.Printf("GenerateInvoiceNumber: %s\n", code)
	PrintFormattedDate()
}

//func GetNumber(number int, Add func(int) int) {
//	result := Add(number)
//	fmt.Println(result)
//}
//
//func Add(a int) int {
//	return a * a
//}
//
//func ProcessNumber(num int, f1 func(int) int, f2 func(int) int) int {
//	result := f1(num)
//	result = f2(result)
//	return result
//}

func Randomizer() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		randomIndex := rand.Intn(len(chars))
		password += string(chars[randomIndex])
	}
	return password
}

//Создайте функцию GenerateInvoiceNumber, которая генерирует и печатает уникальный номер счета в формате "INV-2024-07-001"

func GenerateInvoiceNumber() string {
	var counter int = 0
	now := time.Now()
	counter++
	return fmt.Sprintf("INV-%04d-%02d-%03d", now.Month(), now.Year(), counter)
}

//Создайте функцию PrintFormattedDate, которая печатает текущую дату в формате "2 January 2006"

func PrintFormattedDate() {
	input := time.Now()
	output := input.Format("2 January 2006")
	fmt.Println(output)
}
