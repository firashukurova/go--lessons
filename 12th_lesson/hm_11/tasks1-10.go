package main

import (
	"strings"
	"unicode/utf8"
)

func main() {

}

//1 Напишите функцию, которая принимает две строки и возвращает их конкатенацию

func ConcatenationStrings(str1, str2 string) string {
	return str1 + str2
}

//2 Напишите функцию, которая принимает строку и возвращает её длину.

func ReturnSrtLen(str string) int {
	return utf8.RuneCountInString(str)

}

//3 Напишите функцию, которая проверяет, содержит ли строка заданную подстроку

func CheckSubstr(str, substr string) bool {
	return strings.Contains(str, substr)
}

//4 Напишите функцию, которая находит индекс первого вхождения подстроки в строке. Верните -1, если подстрока не найдена

func FirstInputIdx(str, substr string) int {
	return strings.Index(str, substr)
}

//5 Напишите функцию, которая заменяет все вхождения подстроки в строке на другую подстроку

func ChangeInputSub(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}

//6 Напишите функцию, которая удаляет пробелы в начале и в конце строки

func DeleteSpaces(str string) string {
	return strings.TrimSpace(str)
}
