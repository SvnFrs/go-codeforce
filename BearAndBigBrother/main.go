package main

import (
	"fmt"
)

func main() {
	var limak, bob, year int

	fmt.Scan(&limak, &bob)

	for year = 0; limak <= bob; year++ {
		limak *= 3
		bob *= 2
	}
	fmt.Println(year)
}
