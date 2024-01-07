package main

import "fmt"

func main() {
	var string1, string2 string
	fmt.Scanln(&string1)
	fmt.Scanln(&string2)
	plus := ""
	for i := 0; i < len(string1); i++ {
		if string1[i] == string2[i] {
			plus += "0"
		} else {
			plus += "1"
		}
	}
	fmt.Println(plus)
}
