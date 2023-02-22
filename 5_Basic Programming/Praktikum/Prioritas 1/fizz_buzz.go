package main

import (
	"fmt"
)

func main() {
	data := (100)
	for i := 1; i <= data; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FIZZ_BUZZ")
		} else if i%3 == 0 {
			fmt.Println("FIZZ")
		} else if i%5 == 0 {
			fmt.Println("BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}
