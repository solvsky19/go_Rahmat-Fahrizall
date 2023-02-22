package main

import "fmt"

func main() {
	var kata string
	fmt.Scan(&kata)

	result := ""

	for i := len(kata) - 1; i >= 0; i-- {
		result = result + string(kata[i])
	}
	// fmt.Println(result)
	// fmt.Println(kata)
	if kata == result {
		fmt.Println("palingdrom")
	} else {
		fmt.Println("bukan palingdrom")
	}
}
