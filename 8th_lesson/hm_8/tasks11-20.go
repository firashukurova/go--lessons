package main

import "fmt"

func main() {
	arr := []int{1, -3, 5, 19, 0, 6, -7, 8, 9, -12}
	fmt.Printf("Исходный массив: %v\n", arr)
	fmt.Println("Поменяли индексы:", ChangeArr(arr))

	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(CheckPalindrome(arr1))
	fmt.Println(FindSecondMax(arr1))
	fmt.Println(ReverseArr(arr1))

}

//Поменять местами максимальный и минимальный элементы массива

func ChangeArr(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	minIdx, maxIdx := 0, 0
	min, max := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			minIdx = i
		}
		if arr[i] > max {
			max = arr[i]
			maxIdx = i
		}
	}

	arr[minIdx], arr[maxIdx] = arr[maxIdx], arr[minIdx]
	return arr

}

//Проверить, является ли массив палиндромом

func CheckPalindrome(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

//Найти второе наибольшее число в массиве

func FindSecondMax(arr []int) int {

	max, secondMax := arr[0], arr[1]
	if secondMax > max {
		max, secondMax = secondMax, max
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		} else if arr[i] > secondMax && arr[i] != max {
			secondMax = arr[i]
		}
	}

	return secondMax
}

//Перевернуть массив

func ReverseArr(arr []int) []int {
	reversed := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

//Удалить дубликаты из массива

//func DeleteArr(arr []int) []int {
//
//}
