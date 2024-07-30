package main

import (
	"fmt"
)

func main() {
	//num := 6
	//if num > 9 && num < 100 {
	//	fmt.Println("Двузначное")
	//}
	//
	//n := 6
	//if n%2 == 1 {
	//	fmt.Println("Число нечетное")
	//} else {
	//	fmt.Println("Число четное")
	//}
	//
	////m := 78
	//var m int
	//fmt.Scanln(&m)
	//if m >= 0 && m < 10 {
	//	fmt.Println("Число однозначное")
	//} else if m >= 10 && m < 100 {
	//	fmt.Println("Число двузначное")
	//} else if m >= 100 && m < 1000 {
	//	fmt.Println("Число трехзначное")
	//} else {
	//	fmt.Println("Число слишком большое")
	//}
	//
	//var a, b int
	//fmt.Println("Add A number")
	//fmt.Scanln(&a)
	//fmt.Println("Add B number")
	//fmt.Scanln(&b)
	//if a > b {
	//	fmt.Println("A is bigger than B")
	//} else {
	//	fmt.Println("B is smaller than A")
	//}

	// n1<=n2<=n3
	// n1<=n3<=n2
	// n2<=n1<=n3
	// n2<=n3<=n1
	// n3<=n1<=n2
	// n3<=n2<=n1

	var x, y, z = 9, -19, 15
	if x <= y && y <= z {
		fmt.Println(z, y, x)
	} else if x <= z && z <= y {
		fmt.Println(y, z, x)
	} else if y <= x && x <= z {
		fmt.Println(z, x, y)
	} else if y <= z && z <= x {
		fmt.Println(z, y, x)
	} else if z <= x && x <= y {
		fmt.Println(z, x, y)
	} else if z <= y && y <= x {
		fmt.Println(y, z, x)
	}

	n1 := 1
	switch n1 {
	case 1:
		fmt.Println("Positive")
	case -1:
		fmt.Println("Negative")
	case 0:
		fmt.Println("Zero")
	}

	//найти наибольшее и сложить
	var x1, y1, z1 = 9, -19, 15
	if x1 <= y1 && y1 <= z1 {
		fmt.Println(z1 + y1)
	} else if x1 <= z1 && z1 <= y1 {
		fmt.Println(y1 + z1)
	} else if y1 <= x1 && x1 <= z1 {
		fmt.Println(z1 + x1)
	} else if y1 <= z1 && z1 <= x1 {
		fmt.Println(z1 + y1)
	} else if z1 <= x1 && x1 <= y1 {
		fmt.Println(z1 + x1)
	} else if z1 <= y1 && y1 <= x1 {
		fmt.Println(y1 + z1)
	}

	//узнать високосный год if28
	var year int
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("leap year")

	} else {
		fmt.Println("Not leap year")
	}

	fmt.Println("Add month")
	var month int
	fmt.Scanln(&month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Printf("In month %d is %d days", month, 31)
	case 4, 6, 9, 11:
		fmt.Printf("In month %d is %d days", month, 30)
	case 2:
		fmt.Printf("In month %d is %d days", month, 28)
	}

}
