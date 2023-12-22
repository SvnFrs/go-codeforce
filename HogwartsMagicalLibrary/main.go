package main

import "fmt"

func main() {
	var t, n, x, cnt int64
	fmt.Scan(&t)

	for ; t > 0; t-- {
		fmt.Scan(&n)
		cnt = 0

		for p := int64(0); p < 2*n; p++ {
			fmt.Scan(&x)
			cnt += (x % 2)
		}

		if cnt == n {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
