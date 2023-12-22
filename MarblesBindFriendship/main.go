package main

import (
	"fmt"
	"math"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var boxes int
		fmt.Scan(&boxes)

		marbles := make([]int, boxes)
		totalReductions := 0
		minValue := math.MaxInt

		for j := 0; j < boxes; j++ {
			fmt.Scan(&marbles[j])

			totalReductions += marbles[j]

			if marbles[j] < minValue {
				minValue = marbles[j]
			}
		}

		totalReductions -= minValue * boxes

		fmt.Println(totalReductions)
	}
}
