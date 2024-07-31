package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	target1 := 9
	result1 := FindSubArray(nums1, target1)
	fmt.Printf("Массив %v, цель %d: найдено подмассивов: %d\n", nums1, target1, result1)

}

//21 Найти все подмассивы, сумма которых равна заданному числу

func FindSubArray(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				count++
			}
		}
	}
	return count
}

//22 Найти пару элементов в массиве, сумма которых равна заданному числу

//23 Найти наименьший положительный элемент, отсутствующий в массиве

//24 Найти максимальную сумму подмассива с условием, что подмассив не должен содержать более двух различных элементов

//25 Найти максимальную длину подмассива, сумма элементов которого равна заданному числу

//26 Найти максимальное произведение трех чисел в массиве

//27 Найти подмассив с максимальной суммой

//28 Переместить все отрицательные числа в начало массива, сохраняя порядок остальных чисел

//29 Найти подмассив с наибольшей длиной, сумма элементов которого равна нулю

//30 Найти наибольший общий делитель всех элементов массива
