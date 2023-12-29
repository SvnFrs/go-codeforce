package main

import (
	"fmt"
	"strconv"
)

func onlyContainsFourAndSeven(input string) bool {
	var count int
	for _, char := range input {
		if char == '4' || char == '7' {
			count++
		}
	}
	if count == 4 || count == 7 {
		return true
	} else {
		return false
	}
}

func main() {
	var input int64
	fmt.Scanln(&input)
	strInput := strconv.Itoa(int(input))

	if onlyContainsFourAndSeven(strInput) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
