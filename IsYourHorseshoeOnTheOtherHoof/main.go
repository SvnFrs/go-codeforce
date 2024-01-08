package main

import "fmt"

func main() {
	occured := map[int]bool{}
	var result int
	shoes := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&shoes[i])
	}
	for e := 0; e < 4; e++ {
		if occured[shoes[e]] {
			result++
		} else {
			occured[shoes[e]] = true
		}
	}
	fmt.Println(result)

}
