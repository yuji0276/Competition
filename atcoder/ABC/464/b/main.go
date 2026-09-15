package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(reader, &h, &w)
	grid := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	rmin, rmax := h, -1
	cmin, cmax := w, -1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '#' {
				if i < rmin {
					rmin = i
				}
				if i > rmax {
					rmax = i
				}
				if j < cmin {
					cmin = j
				}
				if j > cmax {
					cmax = j
				}
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := rmin; i <= rmax; i++ {
		fmt.Fprintln(writer, grid[i][cmin:cmax+1])
	}
}
