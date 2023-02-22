package main

import "fmt"

func main() {
	var result int = 5

	for i := 1; i <= result; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(" ")

		}
		for j := 1; j <= i; j++ {
			fmt.Print("# ")
		}
		fmt.Println()
	}

}
