package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	perimeter := a + b + c
	half := perimeter / 2
	area := math.Sqrt(float64(half) * (float64(half) - a) * (float64(half) - b) * (float64(half) - c))
	fmt.Println(perimeter)
	fmt.Println(area)
}
