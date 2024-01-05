package main

import (
	"fmt"
	"strconv"
)

func main() {
	var direction string
	var x, y int

	fmt.Scanln(&direction)

	for _, c := range direction {
		if c == 'N' {
			y++
		}
		if c == 'S' {
			y--
		}
		if c == 'E' {
			x++
		}
		if c == 'W' {
			x--
		}
		if _, err := strconv.Atoi(string(c)); err == nil {
			num := int(c)
			if c+1 == 'N' {
				y += num
			}
			if c+1 == 'S' {
				y -= num
			}
			if c+1 == 'E' {
				x += num
			}
			if c+1 == 'W' {
				x -= num
			}
			c++
		}
	}
	fmt.Println(x, y)
}
