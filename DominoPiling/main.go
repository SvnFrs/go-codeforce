package main

import (
	"fmt"
)

func main() {
	var height int
	var width int

	fmt.Scan(&height)
	fmt.Scan(&width)

	sizes := height * width

	fits := sizes / 2

	fmt.Print(fits)
}
