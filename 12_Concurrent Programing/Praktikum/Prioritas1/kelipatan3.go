// menerapkan goroutine dan channel

package main

import (
	"fmt"
)

func printMultiples(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- 3 * i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go printMultiples(ch)
	for n := range ch {
		fmt.Println(n)
	}
}
