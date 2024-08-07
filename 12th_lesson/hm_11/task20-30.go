package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	s := "You can do it"
	wordCount, charCount := Counter(s)
	fmt.Printf("Words: %d, Characters: %d\n", wordCount, charCount)

	s1 := "you"
	substrings := GenerateSubstrings(s1)
	fmt.Println("All substrings:", substrings)

}

//26 Напишите функцию, которая вычисляет хеш-значение строки с использованием 	алгоритма SHA-256

//27 Напишите функцию, которая генерирует все подстроки заданной строки.

func GenerateSubstrings(str string) []string {
	var result []string
	n := len(str)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			result = append(result, str[i:j])
		}
	}

	return result
}

//28 Напишите функцию, которая ищет все анаграммы строки в другой строке

//29 Напишите функцию, которая подсчитывает количество слов и символов в строке

func Counter(str string) (int, int) {

	charCount := utf8.RuneCountInString(str)
	words := strings.Fields(str)
	wordsLen := len(words)

	return charCount, wordsLen

}
