package main

import "fmt"

func main() {

	var t int
	fmt.Scanf("%d", &t)

	for tc := 0; tc < t; tc++ {

		var d, p int
		fmt.Scan(&d)
		fmt.Scan(&p)

		fmt.Println(d-1, p)
	}
}
