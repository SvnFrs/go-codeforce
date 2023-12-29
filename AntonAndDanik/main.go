package main

import (
	"fmt"
)

func main() {
	var value, Anton, Danik int
	var input string
	fmt.Scanln(&value)
	fmt.Scanln(&input)

	for _, char := range input {
		if char == 'D' {
			Danik++
		}
	}
	Anton = value - Danik

	if Danik > Anton {
		fmt.Println("Danik")
	} else if Danik == Anton {
		fmt.Println("Friendship")
	} else {
		fmt.Println("Anton")
	}
}
