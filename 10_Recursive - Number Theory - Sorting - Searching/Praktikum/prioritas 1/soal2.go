package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostApperItem(items []string) []pair {
	reg := make(map[string]int)
	for _, item := range items {
		reg[item]++
	}

	x := 0
	for _, f := range reg {
		if f > x {
			x = f
		}
	}

	var result []pair
	for item, f := range reg {
		if f == x {
			result = append(result, pair{name: item, count: f})
		}
	}
	return result
}

func main() {
	fmt.Println(MostApperItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostApperItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostApperItem([]string{"football", "basketball", "tenis"}))
}
