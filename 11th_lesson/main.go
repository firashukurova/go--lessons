package main

import (
	"fmt"
	"strings"
)

type StringProcessor struct {
}

func main() {

	s1 := "Hello"
	//s2 := "World"
	fmt.Println(s1[1])

	str := "Hello, world"
	b := str[0]
	fmt.Println(string(b))

	sub := str[7:12]
	fmt.Println(sub)

	strings.Contains(str, "world")

	s := "Hello, 世界"
	bytes := []byte(s)
	fmt.Println(string(bytes))
	fmt.Println(string(bytes))

	for _, v := range s {
		fmt.Println(v)
		fmt.Println(string(v))

	}

}
