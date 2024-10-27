package main

import (
	"fmt"
	"math"
)

func long(radius float64) float64 {
	return math.Pi * radius * radius
}
func main() {
	radius := 3.0
	area := long(radius)
	fmt.Println(area)
}
