package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go printMultiples(5)
	time.Sleep(35 * time.Second)
}
