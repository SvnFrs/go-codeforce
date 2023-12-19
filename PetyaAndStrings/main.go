package main

import (
	"fmt"
	"strings"
)

func main() {
	var first string
	var second string

	fmt.Scan(&first)
	fmt.Scan(&second)

	first = strings.ToLower(first)
	second = strings.ToLower(second)

	result := strings.Compare(first, second)
	fmt.Println(result)

}
