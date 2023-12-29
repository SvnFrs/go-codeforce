package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Scanln(&input)
	hasUpper := 0
	hasLower := 0
	for _, r := range input {
		if unicode.IsUpper(r) {
			hasUpper++
		} else {
			hasLower++
		}
	}
	if hasLower >= hasUpper {
		input = strings.ToLower(input)
	} else {
		input = strings.ToUpper(input)
	}
	fmt.Println(input)

}
