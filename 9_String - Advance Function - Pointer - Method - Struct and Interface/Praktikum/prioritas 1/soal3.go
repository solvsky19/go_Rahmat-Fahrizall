package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {

	if len(a) <= len(b) {
		if strings.Contains(b, a) {
			return a
		}
	} else {
		if strings.Contains(a, b) {
			return b
		}
	}
	return "INVALID"
}
func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}
