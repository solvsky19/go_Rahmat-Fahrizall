package main

import (
	"fmt"
	"strconv"
)

func binaryRepresentation(n int) string {
	if n == 0 {
		return "0"
	}
	var result string
	for n > 0 {
		bit := n % 2
		result = strconv.Itoa(bit) + result
		n = n / 2
	}
	return result
}

func generateBinerArray(n int) []string {
	var ans []string
	for i := 0; i <= n; i++ {
		binary := binaryRepresentation(i)
		ans = append(ans, binary)
	}
	return ans
}

func main() {
	n := 2
	ans := generateBinerArray(n)
	fmt.Println("2", ans) // [0 1 10]

	n = 3
	ans = generateBinerArray(n)
	fmt.Println("3", ans) // [0 1 10 11]
}
