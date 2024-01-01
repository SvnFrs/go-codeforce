package main

import (
	"fmt"
	"strconv"
)

func main() {
	var year int
	fmt.Scanln(&year)

	for ; year > 0; year++ {
		if hasDistinctDigits(year + 1) {
			fmt.Println(year + 1)
			break
		}
	}

}

func hasDistinctDigits(num int) bool {
	digitSet := make(map[rune]bool)

	// Convert the number to a string
	numStr := strconv.Itoa(num)

	// Iterate through each digit in the string
	for _, digit := range numStr {
		// Check if the digit is already in the set
		if digitSet[digit] {
			// If it is, the number has duplicate digits
			return false
		}

		// Otherwise, add the digit to the set
		digitSet[digit] = true
	}

	// If we reach here, all digits are distinct
	return true
}
