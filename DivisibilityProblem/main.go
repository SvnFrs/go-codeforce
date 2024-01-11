package main

import "fmt"

func main() {
	var input, first, second, count int
	fmt.Scan(&input)

	for ; input > 0; input-- {
		count = 0
		fmt.Scanln(&first, &second)
		for first%second != 0 {
			first++
			count++
		}
		fmt.Println(count)
	}
}
