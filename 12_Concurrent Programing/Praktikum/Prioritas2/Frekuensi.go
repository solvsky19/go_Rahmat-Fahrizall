package main

import (
	"fmt"
	"sync"
)

func countFrequency(text string, wg *sync.WaitGroup, mu *sync.Mutex, freq *map[rune]int) {
	defer wg.Done()
	for _, char := range text {
		mu.Lock()
		(*freq)[char]++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	freq := make(map[rune]int)

	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	text2 := "aliqua"

	wg.Add(2)
	go countFrequency(text, &wg, &mu, &freq)
	go countFrequency(text2, &wg, &mu, &freq)

	wg.Wait()

	for char, count := range freq {
		fmt.Printf("%c: %d\n", char, count)
	}
}
