package main

import "fmt"

func main() {
	var number, temp int
	fmt.Scanln(&number)
	value := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scanln(&value[i])
	}
	for i := 1; i < len(value); i++ {
		if value[i] != value[i-1] {
			temp++
		}
	}
	fmt.Println(temp + 1)
}
