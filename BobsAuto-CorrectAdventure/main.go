package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func findAndPrintCharactersToRemove(first, second string) {
	length := int64(len(first))
	var indexesToRemove []int64

	for i := int64(0); i < length; i++ {
		removedString := first[:i] + first[i+1:]

		if removedString == second {
			indexesToRemove = append(indexesToRemove, i+1)
		}
	}

	fmt.Println(len(indexesToRemove))
	if len(indexesToRemove) > 0 {
		sort.Slice(indexesToRemove, func(i, j int) bool {
			return indexesToRemove[i] < indexesToRemove[j]
		})

		result := make([]string, len(indexesToRemove))
		for i, index := range indexesToRemove {
			result[i] = fmt.Sprint(index)
		}
		fmt.Println(strings.Join(result, " "))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var first, second string

	if scanner.Scan() {
		first = scanner.Text()
	} else {
		fmt.Println("0")
		return
	}

	if scanner.Scan() {
		second = scanner.Text()
	} else {
		fmt.Println("0")
		return
	}

	findAndPrintCharactersToRemove(first, second)
}
