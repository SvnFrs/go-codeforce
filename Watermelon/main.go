package main

import (
	"fmt"
)

func main() {
	var watermelon int
	fmt.Scanf("%d", &watermelon)

	if watermelon == 2 || watermelon%2 != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
