package main

import (
	"fmt"
	"strings"
)

func main() {
	var value, Anton, Danik int
	var input string
	fmt.Scanln(&value)
	fmt.Scanln(&input)

	Danik = strings.Count(input, "D")
	Anton = value - Danik
	winner := map[bool]string{Danik > Anton: "Danik", Danik == Anton: "Friendship", Danik < Anton: "Anton"}[true]

	fmt.Println(winner)
}
