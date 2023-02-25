package main

import (
	"fmt"
)

func MunculSekali(angka string) []int {
	key := make(map[int]bool)
	result := make([]int, 0)

	for _, i := range angka {
		num := int(i) - '0'
		if _, j := key[num]; !j {
			key[num] = true
		} else {
			key[num] = false
		}
	}
	for num, nomorUnik := range key {
		if nomorUnik {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("446677"))
}
