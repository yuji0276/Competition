package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(r, &n, &w)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		var ai, bi int
		fmt.Fscan(r, &ai, &bi)
		a[i] = ai
		b[i] = bi
	}

	max := 0
	for i := 0; i < 1<<n; i++ {
		value := 0
		cost := 0
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				value += a[j]
				cost += b[j]
			}
		}

		if (value > max) && (cost <= w) {
			max = value
		}
	}
	fmt.Println(max)
}
