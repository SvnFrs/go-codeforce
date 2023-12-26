package main

import (
	"fmt"
	"math"
)

func main() {
	var distance int
	fmt.Scan(&distance)

	steps := int(math.Ceil(float64(distance) / 5))

	fmt.Println(steps)
}
