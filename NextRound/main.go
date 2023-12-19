package main

import (
	"fmt"
)

func main() {
	var numbers int
	var minimum int
	var contestants []int
	var counter int
	fmt.Scan(&numbers)
	fmt.Scan(&minimum)

	// create an array of contestants
	contestants = make([]int, numbers)
	for i := 0; i < numbers; i++ {
		fmt.Scan(&contestants[i])
	}
	for j := 0; j < numbers; j++ {
		if contestants[j] >= contestants[minimum-1] && contestants[j] != 0 {
			counter++
		}
	}
	fmt.Print(counter)

}
