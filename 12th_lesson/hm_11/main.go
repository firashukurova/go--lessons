package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Hello"
	str2 := "Go"
	//s1 := "Let's"
	//s2 := "programming"
	fmt.Println(Concatenation(str1, str2))
	//fmt.Println(Concatenation2(s1, s2))
	str := "Hello"
	fmt.Println(ReturnLen(str))

}

//1 Напишите функцию, которая принимает две строки и возвращает их конкатенацию.

func Concatenation(str1, str2 string) string {
	con := str1 + " " + str2
	return con
}

//func Concatenation2(str, str2 string) []string {
//	return strings.Split(str, str2)
//}

//2 Напишите функцию, которая принимает строку и возвращает её длину

func ReturnLen(str string) int {
	return len(str)

}

//Дан набор символов (кирилицы). Нужно менять эти символы на следующие

func NextWord(r rune) rune {
	sliseOfLetters := []rune("абвгдеёжзиклмнопрстуфхцчшщъыьэюя")

	for i, letter := range sliseOfLetters {
		if letter == r {
			//если это последняя буква возвращаем первую
			if i == len(sliseOfLetters)-1 {
				return sliseOfLetters[0]
			}
			//иначе возвращаем следующую букву
			return sliseOfLetters[i+1]

		}

	}
	//если не нашли букву в слайсе sliseOfLetters возвращаем без изменений
	return r

}

func ChangeString(str []string) string {
	str = []string{}
	for i := 0; i < len(str); i++ {
		str = append(str, str[i])
	}
	return strings.Join(str, "")

}
