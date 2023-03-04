package main

import "fmt"

var romanNumeralMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func convertToRoman(num int) string {
	result := ""
	for i := range romanNumeralMap {
		for num >= i {
			result += romanNumeralMap[i]
			num -= i
		}
	}
	return result
}

func main() {
	fmt.Println(convertToRoman(4))    // Output: IV
	fmt.Println(convertToRoman(9))    // Output: IX
	fmt.Println(convertToRoman(23))   // Output: XXIII
	fmt.Println(convertToRoman(2021)) // Output: MMXXI
	fmt.Println(convertToRoman(1646)) // Output: MDCXLVI
}
