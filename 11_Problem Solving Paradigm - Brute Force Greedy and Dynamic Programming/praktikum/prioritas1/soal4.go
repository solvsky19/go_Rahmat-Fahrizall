// belum
package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	// Mencari nilai z
	for z := 1; z <= a; z++ {
		if (a-z)*(a-z) > c {
			continue
		}
		for y := 1; y <= a; y++ {
			if z == y {
				continue
			}
			if (a-z-y)*(a-z-y)+(y*y) > c {
				continue
			}
			if z*y*(a-z-y) == b {
				fmt.Printf("x = %d, y = %d, z = %d\n", int(math.Sqrt(float64((a-z-y)*(a-z-y)+(y*y)))), y, z)
				return
			}
		}
	}
	fmt.Println("Tidak ada solusi yang ditemukan.")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
