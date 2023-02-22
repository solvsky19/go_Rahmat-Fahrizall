package main

import "fmt"

func main() {
	number := 12

	fmt.Printf("Faktor dari %d adalah: ", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
