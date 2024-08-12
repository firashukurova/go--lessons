package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// 1 Создайте map для хранения строк и их длин, добавьте несколько элементов и выведите содержимое

	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 10
	m["pear"] = 15
	fmt.Println(m)

	fmt.Println(CheckMapKeys(m, "apple"))
	fmt.Println(CheckMapKeys(m, "orange"))

	UpdateValue(m, "apple", 7)
	fmt.Println(m)

	UpdateValue(m, "cherry", 15)
	fmt.Println(m)

}

//2 Создайте map с несколькими элементами и напишите функцию, которая проверяет наличие заданного ключа

func CheckMapKeys(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}

//3 Напишите функцию для обновления значения в map по заданному ключу

func UpdateValue(m map[string]int, k string, v int) {
	m[k] = v
}

//4 Создайте функцию, которая удаляет элемент из map по заданному ключу

func DeleteElemFromMap(m map[string]int, k string) {
	delete(m, k)
}

//5 Напишите функцию, которая подсчитывает частоту слов в строке и возвращает map с результатами

func CountFreqString(s string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

//6 Напишите функцию, которая сортирует ключи в map и возвращает отсортированные ключи

func SortKeysMap(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

//7 Напишите функцию, которая проверяет, пустой ли map

func IsEmpty(m map[string]int) bool {
	for k := range m {
		if m[k] != 0 {
			return false
		}
	}
	return true
}

//8 Напишите функцию, которая инвертирует ключи и значения в map

func InvertMap(m map[string]int) map[string]int {
	inverted := make(map[string]int)
	for k, v := range m {
		inverted[k] = v
	}
	return inverted
}

//9 Напишите функцию, которая проверяет, есть ли дубликаты значений в map

func FindDuplicateMap(m map[string]int) bool {
	values := make(map[int]bool)
	for _, v := range m {
		if values[v] {
			return false
		}
		values[v] = true

	}
	return false
}

//10 Напишите функцию, которая возвращает все значения из map в срезе

func GetAllValues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

//11 Напишите функцию, которая подсчитывает уникальные значения в срезе строк с использованием map

func FindUniqueValues(slice []string) int {
	uniqueMap := make(map[string]bool)
	for _, v := range slice {
		uniqueMap[v] = true

	}
	return len(uniqueMap)
}

//13 Напишите функцию для объединения двух map. Если ключи совпадают, значения из второго map должны заменять значения из первого.

func UnionMap(m1, m2 map[string]int) map[string]int {
	union := make(map[string]int)
	for k, v := range m1 {
		union[k] = v
	}
	for k, v := range m2 {
		union[k] = v
	}
	return union

}
