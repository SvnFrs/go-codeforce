package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	sum, cnt0 := 0, 0

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
		if a[i] == 0 {
			cnt0++
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	if cnt0 == 0 {
		fmt.Println(-1)
	} else if sum == 0 || a[0] == 0 {
		fmt.Println(0)
	} else {
		if sum%3 == 0 {
			printArray(a)
		} else if sum%3 == 1 {
			removeDigit(a, &sum, 1)
			if sum%3 != 0 {
				removeDigits(a, &sum, 2)
			}
			printNonZeroArray(a)
		} else if sum%3 == 2 {
			removeDigit(a, &sum, 2)
			if sum%3 != 0 {
				removeDigits(a, &sum, 1)
			}
			printNonZeroArray(a)
		}
	}
}

func removeDigit(a []int, sum *int, remainder int) {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != 0 && a[i]%3 == remainder {
			*sum -= a[i]
			a[i] = -1
			break
		}
	}
}

func removeDigits(a []int, sum *int, remainder int) {
	test := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != 0 && a[i]%3 == remainder {
			a[i] = -1
			test++
		}
		if test == 2 {
			break
		}
	}
}

func printArray(a []int) {
	for _, i := range a {
		fmt.Print(i)
	}
}

func printNonZeroArray(a []int) {
	for _, digit := range a {
		if digit != -1 {
			fmt.Print(digit)
		}
	}
}
