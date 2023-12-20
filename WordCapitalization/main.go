package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string

	fmt.Scan(&word)

	capWord := strings.ToUpper(string(word[0]))
	restWord := word[1:]

	fmt.Println(capWord + restWord)
}
