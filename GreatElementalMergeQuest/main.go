package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var n int
		fmt.Scan(&n)

		array := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&array[j])
		}

		sortArray(array)
		result := canReduceToOne(array)

		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func sortArray(arr []int) {
	// Bubble sort for simplicity
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func canReduceToOne(arr []int) bool {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		if arr[i+1]-arr[i] > 1 {
			return false
		}
	}

	return true
}
