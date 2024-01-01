package main

import "fmt"

func main() {
	var people, sum int
	fmt.Scanln(&people)

	think := make([]int, people)

	for i := 0; i < people; i++ {
		fmt.Scanf("%d", &think[i])
		sum += think[i]
	}
	if sum >= 1 {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}
