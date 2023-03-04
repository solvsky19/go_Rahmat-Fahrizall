package main

import "fmt"

type typ struct {
	typeCar string
	fuelIn  int
}

func (e *typ) perkiraanJarak() float64 {
	return float64(e.fuelIn) / (1.5)
}

func main() {
	minibus := typ{
		typeCar: "Minibus",
		fuelIn:  15,
	}

	fmt.Println("Jarak ditempuh:", minibus.perkiraanJarak(), "mill")
}
