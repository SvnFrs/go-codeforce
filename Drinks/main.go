package main

import "fmt"

func main() {
	var n, volume int
	fmt.Scan(&n)
	var totalVolume float64
	for i := 0; i < n; i++ {
		fmt.Scan(&volume)
		totalVolume += float64(volume)
	}
	result := totalVolume / float64(n)
	fmt.Println(result)
}
