package main

import "fmt"

func main() {
	var price, cost, money, banana, borrow int

	fmt.Scan(&price, &money, &banana)

	for i := 1; i <= banana; i++ {
		temp := price * i
		cost += temp
	}
	if cost > money {
		borrow = cost - money
		fmt.Println(borrow)
	} else {
		fmt.Println(0)
	}

}
