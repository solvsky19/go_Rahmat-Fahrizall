package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	for v := 2; v < number; v++ {
		if number%v == 0 {
			return false
		}
	}
	return true
}

// func primeNumber(number int) bool {
// 	var result bool
// 	for i := 1; i < number; i++ {
// 		if number%i == 0 {
// 			fmt.Println(result)
// 		} else {
// 			fmt.Println(result)
// 		}
// 	}
// 	return result
// }

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
