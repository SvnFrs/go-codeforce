package main

import (
	"fmt"
)

func main() {
	// read the binary string
	var binaryString string
	fmt.Scan(&binaryString)

	// create a map to store the mapping between binary strings and digits
	binaryToDigit := make(map[string]string)

	// read the pairwise distinct binary strings 10 characters long representing the numbers 0 to 9
	for i := 0; i < 10; i++ {
		var binaryCode string
		fmt.Scan(&binaryCode)
		binaryToDigit[binaryCode] = fmt.Sprintf("%d", i)
	}

	// initialize the result password string
	var password string

	// iterate over the binary string with steps of 10
	for i := 0; i < len(binaryString); i += 10 {
		// Extract a substring of length 10
		substring := binaryString[i : i+10]

		// map the substring to the original digit and append to the result password
		password += binaryToDigit[substring]
	}

	// print the original password
	fmt.Println(password)
}
