package main

import "fmt"

func main() {
	var testCases, computers, cables, currentComputer, hours int64
	fmt.Scan(&testCases)

	for ; testCases > 0; testCases-- {
		fmt.Scan(&computers, &cables)

		currentComputer, hours = 1, 0

		for currentComputer < computers && currentComputer <= cables {
			currentComputer *= 2
			hours++
		}

		if currentComputer < computers {
			hours += (computers - currentComputer + cables - 1) / cables
		}

		fmt.Println(hours)
	}
}
