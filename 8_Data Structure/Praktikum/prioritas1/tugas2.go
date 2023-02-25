package main

import (
	"fmt"
)

func Mapping(x string, daftar []string) int {
	hasil := 0
	for _, v := range daftar {
		if v == x {
			hasil++
		}
	}
	return hasil
}

func main() {
	x := []string{"agus", "aiman", "sapri", "junet", "aiman", "tuti"}
	y := "aiman"
	result := Mapping(y, x)
	fmt.Printf("%s Muncul sebanyak %d kali dalam daftar\n", y, result)
}
