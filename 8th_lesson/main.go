package main

import "fmt"

func main() {

	//var arr [5]int
	//fmt.Println(arr)

	//var arr2 [5]string
	//fmt.Println(arr2)
	//
	//var matrix [10][3]int
	//fmt.Println(matrix)

	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	//for i, v := range arr {
	//	fmt.Println(i, v)
	//}

	//arr := [5]int{1, 2, 3, 4, 5}
	//slice := arr[2:4] // срез содержит элементы со 2-го по 4-й (индексы 1,2,3)
	//
	//fmt.Println(slice)
	//
	//var emptySlice []int
	//fmt.Println(emptySlice)
	//
	//slice2 := []int{1, 2, 3}
	//slice2 = append(slice, 4, 5, 6)
	//fmt.Println(slice2)

	//sl := make([]int, 5)
	//fmt.Println(sl)
	//sls := make([]int, 5, 10)
	//sls = append(sls, 6, 7, 9)
	//fmt.Println(sls)

	//slice2 := []int{1, 2, 3}
	//slice1 := []int{4, 5}
	//slice3 := append(slice2, slice1...)
	//fmt.Printf("Длина слайса: %d\n", len(slice3))
	//fmt.Printf("Емкость слайса: %d\n", cap(slice3))
	////fmt.Printf("%d\n, %d\n", len(slice3), cap(slice3))
	//
	//src := []int{1, 2, 3}
	//dst := make([]int, len(src))
	//n := copy(dst, src) // копирование всех элементов из src в dst
	//fmt.Println(dst)    // [1 2 3]
	//fmt.Println(n)      // 3

	//nameArray := []string{"F", "i", "r", "u", "z", "a"}
	//nameSlice := nameArray[:3]
	//nameSlice[len(nameSlice)-1] = "R"
	//
	//fmt.Println(nameSlice)
	//fmt.Println(nameArray)

	// Создайте массив из 7 целых чисел и инициализируйте его значениями от 1 до 7.

	numberArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numberArr)

	// Создайте массив из 5 строк и инициализируйте его значениями "a", "b", "c", "d", "e".

	stringArr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(stringArr)

	// Создайте массив из 4 логических значений и инициализируйте его значениями true, false, true, false

	boolArr := [4]bool{true, false, true, false}
	fmt.Println(boolArr)

	// Создайте массив из 10 целых чисел без инициализации и выведите его на экран.

	tenArr := [10]int{}
	fmt.Println(tenArr)

	// Создайте массив из 4 логических значений и выведите значения по индексам 1 и 3.

	blArr := [4]bool{}
	fmt.Println(blArr[1:4])

	// Создайте массив из 3 логических значений и измените значение первого элемента на false

	bl2 := [3]bool{true, false, true}
	bl2[0] = false
	fmt.Println(bl2)

	//Создайте массив из 4 строк и выведите все его элементы с помощью цикла for range.

	strArr := [4]string{"a", "b", "c", "d"}
	for i, v := range strArr {
		fmt.Println(i, v)
	}

	//Дан целочисленный массив A размера 10. Вывести порядковый номер последнего из тех его элементов AK , которые удовлетворяют двойному
	//неравенству A1 < AK < A10. Если таких элементов нет, то вывести 0

	A := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lastIndex := 0
	for i := 1; i < 9; i++ {
		if A[0] < A[i] && A[i] < A[9] {
			lastIndex = i
			//break
		}
	}
	fmt.Println(lastIndex)

	//Дан целочисленный массив размера N, не содержащий одинаковых чисел. Проверить, образуют ли его элементы арифметическую прогрессию
	//(см. задание Array3). Если образуют, то вывести разность прогрессии, если нет — вывести 0

	//n := [5]int{}
	//
	//for i, v := range n {
	//
	//}

}
