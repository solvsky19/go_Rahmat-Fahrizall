package main

import "fmt"

func main() {
	data := 20
	for i := 1; i < data; i++ {
		fmt.Print("Masukan Angka :")
		fmt.Scan(&i)
		if i <= data {
			if i%2 == 0 {
				fmt.Print(i, "Adalah bilangan genap \n")
			} else {
				fmt.Print(i, "Adalah bilangan ganjil\n")
			}
		} else {
			fmt.Print("pernyataan selesai")
		}

	}
}
