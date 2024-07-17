package main

import "fmt"

func main() {
	CheckNumber()
	CheckAge()
	CheckPassword()
}

//Создайте функцию CheckNumber, которая принимает целое число и печатает "Положительное", если число больше нуля; "Отрицательное",
//если меньше нуля; и "Ноль", если равно нулю

func CheckNumber() {
	var number int
	fmt.Println("Введите число")
	fmt.Scan(&number)
	if number > 0 {
		fmt.Println("Положительное")
	} else if number == 0 {
		fmt.Println("Ноль")
	} else {
		fmt.Println("Отрицательное")
	}
}

//Создайте функцию CheckAge, которая принимает возраст и печатает "Совершеннолетний", если возраст 18 и старше; "Несовершеннолетний", если младше 18

func CheckAge() {
	var age int
	fmt.Println("Введите возраст")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}

}

//Создайте функцию CheckPassword, которая принимает строку пароля и печатает "Пароль верный", если пароль равен "1234"; и "Пароль неверный" в противном случае

func CheckPassword() {
	var password int
	fmt.Println("Введите пароль в цифрах")
	fmt.Scan(&password)
	if password == 1234 {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверен, повторите снова")
	}
	//fmt.Scanln(&password)
}
