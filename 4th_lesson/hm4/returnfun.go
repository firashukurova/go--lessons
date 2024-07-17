package main

import "fmt"

func main() {
	score := GetGrade()
	fmt.Println(score)

	amount := GetDiscount()
	fmt.Println(amount)

	temperature := GetTemperatureDescription()
	fmt.Println(temperature)
}

//Создайте функцию GetGrade, которая возвращает оценку "A", "B" или "C" в зависимости от значения переменной score.
//Переменную scope вводить с консоли внутри функции

func GetGrade() string {
	var score int
	fmt.Println("Введите оценку (3, 4 или 5):")
	fmt.Scan(&score)
	switch score {
	case 5:
		return "A"
	case 4:
		return "B"
	case 3:
		return "C"
	default:
		return "Некорректная оценка"
	}
}

//Создайте функцию GetDiscount, которая возвращает скидку "10%" - при amount > 1000, "5%" - при 500 < amount <=100
//или "0%" - при amount <= 500 в зависимости от суммы покупки amount. Переменную amount вводить с консоли внутри функции

func GetDiscount() string {
	var amount int
	fmt.Println("Введите сумму покупки:")
	fmt.Scan(&amount)
	if amount > 1000 {
		return "10%"
	} else if amount > 500 && amount <= 1000 {
		return "5%"
	} else {
		return "0%"
	}
}

//Создайте функцию GetTemperatureDescription, которая возвращает "Холодно" (меньше 10) , "Тепло" (с 10 до 25)  или
//"Жарко" в зависимости от значения переменной temperature. Переменную temperature вводить с консоли внутри функции

func GetTemperatureDescription() string {
	var temperature int
	fmt.Println("Введите температуру в °C")
	fmt.Scan(&temperature)
	if temperature < 10 {
		return "Холодно"
	} else if temperature > 10 && temperature <= 25 {
		return "Тепло"
	} else {
		return "Жарко"
	}
}
