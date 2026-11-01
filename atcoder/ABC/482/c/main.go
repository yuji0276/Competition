package main

import "fmt"

func main() {
	var h, w, n, q int
	fmt.Scan(&h, &w, &n, &q)

	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}
	for range n {
		var r, c int
		fmt.Scan(&r, &c)
		grid[r-1][c-1] = 1
	}

	for range q {
		var s, u, d int
		fmt.Scan(&s, &u, &d)
		if s == 1 {
		} else {
		}
	}
}
