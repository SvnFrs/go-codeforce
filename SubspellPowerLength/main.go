package main

import (
	"fmt"
)

func maxLengthSubsequence(n int, a []int) (int, []int) {
	// Check and modify the array to ensure a[i+1] > a[i] or a[i-1] < a[i]
	for i := 1; i < n-1; i++ {
		if a[i] <= a[i-1] && a[i] <= a[i+1] {
			// Modify the current element to make it strictly increasing
			a[i]++
		} else if a[i] >= a[i-1] && a[i] >= a[i+1] {
			// Modify the current element to make it strictly decreasing
			a[i]--
		}
	}

	// Check the length of the strictly increasing sequence
	length := 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			length++
		}
	}

	return length, a
}

func main() {
	// Read input values
	var n int
	fmt.Print("Enter the size: ")
	fmt.Scan(&n)

	a := make([]int, n)
	fmt.Print("Enter the array elements: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	// Find and print the maximum length of the required subsequence
	resultLength, resultArray := maxLengthSubsequence(n, a)
	fmt.Printf("Length of the increasing subarray: %d\n", resultLength)
	fmt.Printf("Modified Array: %v\n", resultArray)
}
