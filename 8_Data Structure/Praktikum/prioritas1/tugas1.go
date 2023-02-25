package main

import (
	"fmt"
)

func ArrayMerge(array1 []string, array2 []string) []string {
	result := append(array1, array2...)
	tampung := make(map[string]bool)
	x := make([]string, 0)

	for _, v := range result {
		if _, i := tampung[v]; !i {
			tampung[v] = true
			x = append(x, v)
		}
	}

	return x
}

func main() {
	array1 := []string{"aa", "ai", "au", "ae", "ao"}
	array2 := []string{"ra", "ai", "ru", "re", "ae"}
	result := ArrayMerge(array1, array2)

	fmt.Println(result)
}
