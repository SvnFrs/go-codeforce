package main

import "fmt"

func main() {
	var n, result int
	fmt.Scan(&n)

	if n%2 == 0 {
		result = n / 2
	} else {
		result = -(n + 1) / 2
	}

	fmt.Println(result)
}
