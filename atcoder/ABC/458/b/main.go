package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var H, W int

	fmt.Fscan(reader, &H, &W)

	grid := make([][]int, H)
	for i := range grid {
		grid[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if (H == 1) && (W == 1) {
				grid[0][0] = 0
				break
			}
			if (i == 0) || (i == H-1) {
				grid[i][j] = 3
				grid[i][0] = 2
				grid[i][W-1] = 2
			} else {
				grid[i][j] = 4
				grid[i][0] = 3
				grid[i][W-1] = 3
			}
			if (H == 1) || (W == 1) {
				grid[i][j] = 2
				grid[0][0] = 1
				grid[0][W-1] = 1
				grid[H-1][0] = 1
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Printf("%v ", grid[i][j])
		}
		fmt.Println()
	}
}
