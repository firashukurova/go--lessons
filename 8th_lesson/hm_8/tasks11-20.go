package main

import "fmt"

func main1() {
	arr := []int{1, -3, 5, 19, 0, 6, -7, 8, 9, -12}
	fmt.Printf("Исходный массив: %v\n", arr)
	fmt.Println("Поменяли индексы:", ChangeArr(arr))

	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(CheckPalindrome(arr1))
	fmt.Println(FindSecondMax(arr1))
	fmt.Println(ReverseArr(arr1))

	arr3 := []int{1, 3, 3, 4, 4, 5, 5, 6, 7, 8, 9}
	fmt.Println(DeleteDuplicate(arr3))

	zeroArr := []int{0, 0, 7, 3, 0, 1, 0, 12, 3}
	fmt.Println(MoveZero(zeroArr))

	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a2 := []int{9, 44, 30, 12, 5, 7, 15}
	fmt.Println(IntersectionArr(a1, a2))
	fmt.Println(SubsetOfArray(a1, a2))

}

//11Поменять местами максимальный и минимальный элементы массива
//необходимо добавить если будет 2 или 3 макс и мин

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

//12 Проверить, является ли массив палиндромом

func CheckPalindrome(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

//13 Найти второе наибольшее число в массиве

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

//14 Перевернуть массив

func ReverseArr(arr []int) []int {
	reversed := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed

}

//15 Удалить дубликаты из массива

func DeleteDuplicate(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	res := []int{}
	for i, v := range arr {
		found := false
		for j := 0; j < i; j++ {
			if arr[j] == v {
				found = true
				break
			}
		}
		if !found {
			res = append(res, v)
		}
	}
	return res
}

//16 Переместить все нули в конце массива, сохраняя порядок ненулевых элементов

func MoveZero(arr []int) []int {
	nonZeroIndex := 0
	for _, v := range arr {
		if v != 0 {
			arr[nonZeroIndex] = v
			nonZeroIndex++
		}

	}
	for i := nonZeroIndex; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr

}

//17 Найти пересечение двух массивов

func IntersectionArr(arr1, arr2 []int) []int {
	intersection := []int{}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				intersection = append(intersection, arr1[i])
			}
		}
	}
	return intersection
}

//18 Проверить, является ли массив подмножеством другого массива

func SubsetOfArray(arr1, arr2 []int) bool {
	for _, elem1 := range arr1 {
		found := false
		for _, elem2 := range arr2 {
			if elem1 == elem2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

//19 Объединить два отсортированных массива в один отсортированный

func UnionSortedArr(arr1, arr2 []int) []int {
	result := []int{}

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	//Добавим оставшиеся из arr1 элементы если они есть
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	//Добавим оставшиеся из arr2 элементы если они есть
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}

//20 Найти длину самого длинного подмассива, в котором все элементы различны

func LongestUniqueSubarray(arr []int) int {
	maxLength := 0
	start := 0

	for end := 0; end < len(arr); end++ {
		// Проверяем, есть ли arr[end] в текущем окне arr[start:end]
		for i := start; i < end; i++ {
			if arr[i] == arr[end] {
				// Если нашли повторение, сдвигаем start
				start = i + 1
				break
			}
		}

		// Обновляем maxLength, если текущее окно больше
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength

}

//15 Удалить дубликаты из массива

//func DeleteArr(arr []int) []int {
//	res := []int{}
//	for i, v := range arr {
//		found := false
//		for j := i + 1; j < len(arr); j++ {
//			if arr[j] == v {
//				found = true
//			}
//		}
//	}
//	res = append(res, arr[len(arr)-1])
//	return res
//}
