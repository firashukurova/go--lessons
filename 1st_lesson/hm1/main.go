package main

import (
	"fmt"
	//"strconv"
)

// 1) Объявите переменные различных целочисленных типов данных (int, int8, int16, int32, int64),
// вещественных типов данных (float32 и float64), логический тип данных(bool), rune, byte, строки(string), uint-ы, комплексные числа(complex64, complex128)
// и узнать какое у них значение по умолчанию. Напишите программу, которая по ходу меняет значения этих переменных и выводит на экран их значения.
// (Можете использовать вашу фантазию)
// 2) Объявить 2 переменные с типом данных int и произвести с ними арифметические операции

func main() {

	var int1 int
	var int2 int8
	var int3 int16
	var int4 int32
	var int5 int64

	var uint1 uint
	var uint2 uint8
	var uint3 uint32
	var uint4 uint64

	var float1 float32
	var float2 float64

	var bl bool

	var rn rune

	var bt byte

	var str string

	var cmpx complex64
	var cmpx2 complex128

	fmt.Print(int1, "\n", int2, "\n", int3, "\n", int4, "\n", int5, "\n")
	fmt.Print(uint1, "\n", uint2, "\n", uint3, "\n", uint4, "\n")

	fmt.Print(float1, "\n", float2, "\n")

	fmt.Print(bl, "\n", rn, "\n", bt, "\n", str, "\n", cmpx, "\n", cmpx2, "\n")

	int1 = 10
	int2 = -123
	int3 = -321
	int4 = 1009
	int5 = 2900

	uint1 = 10
	uint2 = 98
	uint3 = 987
	uint4 = 2983

	float1 = 3.9
	float2 = 9.5

	bl = true

	rn = 'A'

	bt = 255

	str = "Welcome"

	cmpx = 2 + 3i
	cmpx2 = 5.5 + 5i

	fmt.Print(int1, "\n", int2, "\n", int3, "\n", int4, "\n", int5, "\n")
	fmt.Print(uint1, "\n", uint2, "\n", uint3, "\n", uint4, "\n")

	fmt.Print(float1, "\n", float2, "\n")

	fmt.Print(bl, "\n", rn, "\n", bt, "\n", str, "\n", cmpx, "\n", cmpx2, "\n")

	//второе задание
	num1 := 30
	num2 := 50

	fmt.Println(num1 + num2)

	a := 54
	b := 89
	c := a + b

	fmt.Println(c)

}
