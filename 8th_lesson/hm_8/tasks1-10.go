package main

import "fmt"

func main() {

	array := []int{1, 2, -8, 4, 5, -3, 7, -1, 10, 12}

	min, max := FindMinMax(array)
	fmt.Printf("Минимальный элемент: %d\n", min)
	fmt.Printf("Максимальный элемент: %d\n", max)

	slice := CopyArr(array)
	fmt.Printf("Исходный массив: %v\n", array)
	fmt.Printf("Скопированный массив: %v\n", slice)

	arr1 := []int{1, -3, 7, -5, 13, 5, -2, 10, 9}

	positive := FindPositive(arr1)
	fmt.Printf("Positive numbers: %d\n", positive)

	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := FindSum(arr2)
	fmt.Printf("Сумма всех элементов: %d\n", sum)
	fmt.Printf("Среднее значение элементов массива:%2.f\n", FindAvg(arr2))

	arr3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numToDelete := 5
	result := DeleteArr(arr3, numToDelete)
	fmt.Println(result)

	arr := []int{1, 2, 3, 4, 5}
	multiplier := 2
	result = MultiplyArr(arr, multiplier)
	fmt.Println(result)

	idxArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 3
	idx := FindIdx(idxArr, num)
	fmt.Println(idx)

	sliceU := []string{"Hello"}
	sliceU2 := []string{"Go"}
	resultUnion := UnionArr(sliceU, sliceU2)
	fmt.Println(resultUnion)

}

//1 и 2) Найти максимальный элемент в массиве. Найти минимальный элемент в массиве.

func FindMinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

//3 Подсчитать количество положительных чисел в массиве.

func FindPositive(arr []int) int {
	counter := 0
	for _, v := range arr {
		if v > 0 {
			counter++
		}
	}
	return counter
}

//4 Найти сумму всех элементов массива.

func FindSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

//5 Найти среднее значение всех элементов массива.

func FindAvg(arr []int) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += float64(v)

	}
	avg := sum / float64(len(arr))
	return avg
}

//6 Удалить все вхождения заданного числа из массива

func DeleteArr(arr []int, num int) []int {
	result := make([]int, 0, len(arr))
	for _, v := range arr {
		if v != num {
			result = append(result, v)
		}

	}
	return result

}

//7 Умножить все элементы массива на заданное число

func MultiplyArr(arr []int, multiplier int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = v * multiplier
	}
	return result
}

//8 Найти все индексы заданного числа в массиве

func FindIdx(arr []int, target int) []int {
	result := make([]int, 0)
	for i, v := range arr {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}

//9 Создать копию массива

func CopyArr(arr []int) []int {
	slice := []int{} //make([]int, len(arr))
	copy(slice, arr)
	return slice
}

//10 Объединить два массива

func UnionArr(arr1, arr2 []string) []string {
	combinedSlice := append(arr1, arr2...)
	return combinedSlice

}
