package main

import "fmt"

func main() {
	var nilai = 12

	if nilai <= 100 && nilai >= 80 {
		fmt.Println("nilai A")
	} else if nilai <= 79 && nilai > 65 {
		fmt.Println("nilai B")
	} else if nilai <= 64 && nilai >= 50 {
		fmt.Println("nilai C")
	} else if nilai <= 49 && nilai >= 35 {
		fmt.Println("nilai D")
	} else if nilai <= 34 && nilai >= 0 {
		fmt.Println("nilai E")
	} else {
		fmt.Println("Invalid")
	}
}
