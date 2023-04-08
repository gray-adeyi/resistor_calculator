package main

import (
	"fmt"

	"github.com/gray-adeyi/resistor_calculator/lib"
)

func main() {
	myResistor := lib.FourBandResistor{Band1: lib.Green, Band2: lib.Blue, Band3: lib.Yellow, Tolerance: lib.Silver}
	fmt.Printf("My resistor is %d ohms\n", myResistor.IdealValue())
	min, max := myResistor.ValueRange()
	fmt.Printf("My resistor range is %.2f %.2f ohms\n", min, max)
}
