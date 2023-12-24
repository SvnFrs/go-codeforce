package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	// the map have type rune for the value of the character, then the counts is stored as an integer
	charCount := make(map[rune]int)

	var result string

	for _, char := range input {
		charCount[char]++
		if charCount[char] == 1 {
			result += string(char)
		}
	}
	if len(result)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
