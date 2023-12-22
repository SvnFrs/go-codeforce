package main

import "fmt"

func main() {
	var testCases, floor, hours int64
	fmt.Scan(&testCases)

	for testCases > 0 {
		fmt.Scan(&floor, &hours)
		hours = hours - 1

		if floor%2 == 1 {
			hours += hours / (floor / 2)
		}

		fmt.Println(1 + (hours % floor))
		testCases--
	}
}
