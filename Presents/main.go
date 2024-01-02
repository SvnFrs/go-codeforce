package main

import (
	"fmt"
	"strings"
)

func main() {
	var quantity, index int
	fmt.Scan(&quantity)
	friend := make([]int, quantity)

	for i := 0; i < quantity; i++ {
		fmt.Scan(&index)
		friend[index-1] = i + 1

	}
	result := strings.Trim(fmt.Sprint(friend), "[]")
	fmt.Println(result)
}
