package main

import (
	"fmt"
	"strings"
)

func main() {
	var init = 0
	var plus = "++"
	var minus = "--"
	var input int
	var statements string

	fmt.Scan(&input)

	for i := 0; i < input; i++ {
		fmt.Scan(&statements)
		if strings.Contains(statements, plus) {
			init++
		} else if strings.Contains(statements, minus) {
			init--
		}
	}
	fmt.Print(init)

}
