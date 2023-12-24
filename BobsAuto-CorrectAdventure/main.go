package main

import (
	"fmt"
)

type pi struct {
	f, s int
}

type pl struct {
	f, s int64
}

type pd struct {
	f, s float64
}

func main() {
	solve()
}

func solve() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	ans := -1

	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			ans = i
			break
		}
	}

	if ans == -1 {
		ans = len(s1) - 1
	}

	for j := ans + 1; j < len(s1); j++ {
		if s1[j] != s2[j-1] {
			if ans == -1 {
				ans = j
			} else {
				fmt.Println(0)
				return
			}
		}
	}

	temp := ans
	count := 1
	for temp+1 < len(s1) && s1[temp+1] == s1[temp] {
		count++
		temp++
	}
	for ans-1 >= 0 && s1[ans-1] == s1[ans] {
		count++
		ans--
	}
	fmt.Println(count)

	for i := ans; i <= temp; i++ {
		fmt.Print(i+1, " ")
	}
}

// The equivalent of C++ set is not directly available in Go, so the set-related definitions and functions are not translated.
