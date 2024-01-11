package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		if i%2 == 0 {
			fmt.Printf("I love ")
		} else {
			fmt.Printf("I hate ")
		}
		if i != input {
			fmt.Printf("that ")
		}
	}
	fmt.Printf("it\n")
}
