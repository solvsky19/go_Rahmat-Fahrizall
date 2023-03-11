// goroutine dan buffer channel

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
	ch := make(chan int, 5) // membuat buffer channel dengan kapasitas 5
	go printMultiples(ch)
	for n := range ch {
		fmt.Println(n)
	}
}
