package main

import "fmt"

var (
	grid  [][]int
	n, m  int
	ans   int
	count int
)

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--

		readInput()
		ans = 0

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] != 0 {
					count = 0
					dfs(i, j)
					ans = max(count, ans)
				}
			}
		}

		fmt.Println(ans)
	}
}

func readInput() {
	fmt.Scan(&n, &m)
	grid = make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
		for j := range grid[i] {
			fmt.Scan(&grid[i][j])
		}
	}
}

func dfs(i, j int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 {
		return
	}

	count += grid[i][j]
	grid[i][j] = 0
	dfs(i+1, j)
	dfs(i-1, j)
	dfs(i, j+1)
	dfs(i, j-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
