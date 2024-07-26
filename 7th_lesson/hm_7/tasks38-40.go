package main

import "fmt"

func main() {

	//Создайте переменную анонимного типа, которая содержит Title, Date и Location (анонимная структура с полями Street и City) для события.
	//Напишите функцию, которая принимает эту переменную и выводит информацию о событии

	eventInfo := struct {
		Title    string
		Date     string
		Location struct {
			Street, City string
		}
	}{Title: "IT-Gap",
		Date: "2024-07-25",
		Location: struct {
			Street, City string
		}{
			Street: "проспект Исмоила Сомони, 47",
			City:   "Dushanbe",
		},
	}
	printEventInfo(eventInfo)

	//Создайте переменную анонимного типа, которая содержит Title и Author для книги. Напишите функцию, которая принимает
	//эту переменную и выводит информацию о книге

	bookInfo := struct {
		Title  string
		Author string
	}{
		Title:  "Master & Margarita",
		Author: "M.Bulgakov",
	}
	printBookInfo(bookInfo)

	//Создайте переменную анонимного типа, которая содержит Make, Model и Year для автомобиля. Напишите функцию, которая
	//принимает эту переменную и выводит информацию о автомобиле.

	carInfo := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Germany",
		Model: "BMW",
		Year:  2020,
	}
	printCarInfo(carInfo)
}

func printEventInfo(event struct {
	Title    string
	Date     string
	Location struct {
		Street string
		City   string
	}
}) {
	fmt.Printf("Событие: %s\n", event.Title)
	fmt.Printf("Дата: %s\n", event.Date)
	fmt.Printf("Место: %s, %s\n", event.Location.Street, event.Location.City)
}

func printBookInfo(book struct {
	Title  string
	Author string
}) {
	fmt.Printf("Книга: %s\n", book.Title)
	fmt.Printf("Авто: %s\n", book.Author)
}

func printCarInfo(car struct {
	Make  string
	Model string
	Year  int
}) {
	fmt.Printf("Производитель: %s\n", car.Make)
	fmt.Printf("Марка: %s\n", car.Model)
	fmt.Printf("Год: %d\n", car.Year)
}
