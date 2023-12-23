package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var addition string
	fmt.Scan(&addition)
	if len(addition) < 2 {
		fmt.Println(addition)
	} else {
		parts := strings.Split(addition, "+")
		var numbers []int
		final := ""

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error converting %s to integer: %v\n", part, err)
			} else {
				numbers = append(numbers, num)
			}
		}
		sort.Ints(numbers)
		for _, v := range numbers {
			temp := strconv.Itoa(v) + "+"
			final = final + temp
		}
		fmt.Println(strings.TrimRight(final, "+"))
	}

}
