package main

import "fmt"

func main() {
	var number, people int
	var in, out, seat int
	seat = 0
	fmt.Scan(&number)
	for ; number > 0; number-- {
		fmt.Scanln(&out, &in)
		people -= out
		people += in
		if seat < people {
			seat += people - seat
		}
	}
	fmt.Println(seat)

}
