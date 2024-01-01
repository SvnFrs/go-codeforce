package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var prevMagnet, currentMagnet string
	groupCount := 0

	for scanner.Scan() {
		currentMagnet = scanner.Text()
		if prevMagnet != currentMagnet {
			groupCount++
		}
		prevMagnet = currentMagnet
	}

	// since we are counting transitions, there is one more group than the number of transitions.
	// this handles the case where the last group extends to the end.
	fmt.Println(groupCount - 1)
}
