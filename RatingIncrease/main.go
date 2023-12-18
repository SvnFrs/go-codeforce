package main

import (
	"fmt"
	"strconv"
)

func main() {
	var testcases int
	var input string
	var inputArray []int
	var a int
	var b int
	fmt.Scan(&testcases)

	for i := 0; i < testcases; i++ {
		fmt.Scan(&input)
		inputArray = make([]int, len(input))
		for i := 0; i < len(input); i++ {
			inputArray[i], _ = strconv.Atoi(string(input[i]))
		}
		for i := 0; i < len(inputArray)-1; i++ {
			a = 0
			b = 0
			for j := 0; j <= i; j++ {
				a = a*10 + inputArray[j]
			}
			for j := i + 1; j < len(inputArray); j++ {
				b = b*10 + inputArray[j]
			}
			if b > a && inputArray[i+1] != 0 {
				fmt.Println(a, b)
				break
			}
			if i == len(inputArray)-2 {
				fmt.Println("-1")
			}
		}
	}
}
