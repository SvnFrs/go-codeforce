package main

import "fmt"

func main() {
	var n, t int
	fmt.Scanln(&n, &t)

	var s string
	fmt.Scanln(&s)

	// convert string to byte slice for easy swapping
	queue := []byte(s)

	for time := 1; time <= t; time++ {
		for i := 0; i < n-1; i++ {
			// if the current position has a boy and the next has a girl, swap them
			if queue[i] == 'B' && queue[i+1] == 'G' {
				queue[i], queue[i+1] = queue[i+1], queue[i]
				i++ // skip the next position since it has already been processed in this step
			}
		}
	}

	// convert byte slice back to string
	result := string(queue)

	fmt.Println(result)
}
