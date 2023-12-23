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
		return
	}
	parts := strings.Split(addition, "+")
	numbers := make([]int, len(parts))

	for i, part := range parts {
		if num, err := strconv.Atoi(part); err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", part, err)
			return
		} else {
			numbers[i] = num
		}
	}
	sort.Ints(numbers)
	final := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), "+"), "[]")
	fmt.Println(final)
}
