package main

import "fmt"

func main() {
	var string1, string2, temp string
	fmt.Scanln(&string1)
	fmt.Scanln(&string2)

	temp = reverse(string2)

	if string1 == temp {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
