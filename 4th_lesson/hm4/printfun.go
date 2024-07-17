package main

import "fmt"

func main() {
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
}

// Создайте функцию PrintGreeting, которая печатает "Доброе утро!", если dayType равен "утро"; "Добрый день!",
// если dayType равен "день"; и "Добрый вечер!", если dayType равен "вечер". Переменную  dayType вводить с консоли внутри функции

func PrintGreeting() {
	var dayType string
	fmt.Println("Add day type: morning, afternoon, evening")
	fmt.Scanln(&dayType)
	switch dayType {
	case "morning":
		fmt.Println("Good Morning!")
	case "evening":
		fmt.Println("Good Evening!")
	case "afternoon":
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Incorrect type!")
	}
}

//Создайте функцию PrintWeather, которая печатает "Солнечно", если weatherType равен "солнечно"; "Облачно", если weatherType равен "облачно";
//и "Дождливо", если weatherType равен "дождливо". Переменную  weatherType вводить с консоли внутри функции

func PrintWeather() {
	var weatherType string
	fmt.Println("Add weather type: sunny, cloudy, rainy")
	fmt.Scanln(&weatherType)
	switch weatherType {
	case "sunny":
		fmt.Println("Weather is sunny")
	case "cloudy":
		fmt.Println("Weather is cloudy")
	case "rainy":
		fmt.Println("Weather is rainy")
	default:
		fmt.Println("Incorrect weather type!")
	}
}

//Создайте функцию PrintTrafficLight, которая печатает "Стоп", если lightColor равен "красный"; "Внимание", если lightColor равен "желтый";
//и "Идите", если lightColor равен "зеленый". Переменную lightColor вводить с консоли внутри функции

func PrintTrafficLight() {
	var lightColor string
	fmt.Println("Add light color: red, yellow, green")
	fmt.Scanln(&lightColor)
	switch lightColor {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Attention")
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Incorrect type!")
	}
}
