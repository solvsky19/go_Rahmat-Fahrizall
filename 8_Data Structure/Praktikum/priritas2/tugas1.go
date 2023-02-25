package main

import (
	"fmt"
)

func PairSum(array []int, target int) (int, int) {
	y := 0
	x := len(array) - 1

	for y < x {
		sum := array[y] + array[x]
		if sum == target {
			return y, x
		} else if sum < target {
			y++
		} else {
			x--
		}
	}
	return -1, -1
}
func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 10, 11}, 11))
}
