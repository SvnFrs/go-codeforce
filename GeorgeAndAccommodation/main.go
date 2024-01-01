package main

import "fmt"

func main() {
	var room, people, capacity, YES int
	fmt.Scanln(&room)
	for i := 0; i < room; i++ {
		fmt.Scanln(&people, &capacity)

		if capacity-people > 1 {
			YES++
		}
	}
	fmt.Println(YES)
}
