package main

import "fmt"

//Создайте структуру Course с полями Title и Instructor (структура Person). Напишите функцию, которая принимает Course и выводит информацию о курсе

type Person struct {
	FullName string
	Age      int
	Email    string
}

type Course struct {
	Title      string
	Instructor Person
}

//Создайте структуру Event с полями Title и Location (структура Address). Напишите функцию, которая принимает Event и выводит информацию о событии.

func main() {
	course := Course{
		Title: "Go Programming Language",
		Instructor: Person{
			FullName: "Ergashev Abdukhalim",
			Age:      20,
			Email:    "ergashev@gmail.com",
		},
	}
	CourseInfo(course)
}

func CourseInfo(course Course) {
	fmt.Println("Course Title:", course.Title)
	fmt.Println("Course Instructor:", course.Instructor.FullName)
	fmt.Println("Course Instructor:", course.Instructor.Age)
	fmt.Println("Email:", course.Instructor.Email)
}
