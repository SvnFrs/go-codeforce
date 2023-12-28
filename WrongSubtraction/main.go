package main

import "fmt"

func main() {
	var number, subtraction int
	fmt.Scanln(&number, &subtraction)

	for ; subtraction > 0; subtraction-- {
		if number%10 != 0 {
			number--
		} else {
			number /= 10
		}
	}
	fmt.Println(number)
}
