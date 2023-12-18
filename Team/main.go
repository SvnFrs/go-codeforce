package main

import (
	"fmt"
)

func main() {
	var problems int
	fmt.Scan(&problems)
	views := make([]int, 3)
	var points int
	for i := 0; i < problems; i++ {
		// scan the views
		for j := 0; j < 3; j++ {
			fmt.Scan(&views[j])
		}
		if views[0]+views[1]+views[2] >= 2 {
			points++
		}
	}
	fmt.Println(points)
}
