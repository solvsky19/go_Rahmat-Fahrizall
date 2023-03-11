package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex, result *int) {
	defer wg.Done()
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	mu.Lock()
	*result *= f
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := 1
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go factorial(i, &wg, &mu, &result)
	}
	wg.Wait()
	fmt.Println(result)
}
