package main

import (
	"fmt"
)

func canPassAllLevels(n int, xLevels []int, yLevels []int) string {
	allLevels := make(map[int]bool)
	for i := 1; i <= n; i++ {
		allLevels[i] = true
	}

	// Mark levels that Little X can pass
	xSet := make(map[int]bool)
	for _, level := range xLevels[1:] {
		xSet[level] = true
	}

	// Mark levels that Little Y can pass
	ySet := make(map[int]bool)
	for _, level := range yLevels[1:] {
		ySet[level] = true
	}

	// Check if the union of Little X and Little Y levels covers all levels
	for level := range allLevels {
		if !xSet[level] && !ySet[level] {
			return "Oh, my keyboard!"
		}
	}

	return "I become the guy."
}

func main() {
	var n, p, q int
	fmt.Scan(&n)

	var xLevels []int
	fmt.Scan(&p)
	xLevels = make([]int, p+1)
	for i := 1; i <= p; i++ {
		fmt.Scan(&xLevels[i])
	}

	var yLevels []int
	fmt.Scan(&q)
	yLevels = make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Scan(&yLevels[i])
	}

	result := canPassAllLevels(n, xLevels, yLevels)
	fmt.Println(result)
}
