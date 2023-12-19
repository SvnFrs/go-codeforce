package main

import (
	"fmt"
	"math"
)

func main() {
	// init a 5x5 matrix
	var matrix [5][5]int

	// scan the matrix
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	// calculate the moves to the center
	var moves int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// if the value is 1
			if matrix[i][j] == 1 {
				// calculate the moves to the center
				moves = int(math.Abs(float64(i-2)) + math.Abs(float64(j-2)))
			}
		}
	}
	fmt.Print(moves)

}
