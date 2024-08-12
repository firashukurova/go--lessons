package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 10
	fmt.Println(len(m))

	original := map[string]int{"key1": 10, "key2": 20}
	copy := make(map[string]int)
	for k, v := range original {
		copy[k] = v
	}
	//copy := original
	fmt.Println(copy)
}
