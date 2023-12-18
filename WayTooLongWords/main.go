package main

import (
	"fmt"
)

func main() {
	var lines int
	fmt.Scanf("%d", &lines)

	// fmt.Println(length)
	for i := 0; i < lines; i++ {
		var words string
		fmt.Scan(&words)

		length := len(words)
		wordsArray := make([]string, length)
		for i := 0; i < length; i++ {
			wordsArray[i] = string(words[i])
		}
		// fmt.Println(wordsArray)
		if length > 10 {
			first := wordsArray[0]
			last := wordsArray[len(wordsArray)-1]
			middle := len(wordsArray) - 2
			fmt.Print(first)
			fmt.Print(middle)
			fmt.Print(last + "\n")
		} else {
			fmt.Println(words)
		}
	}
}
