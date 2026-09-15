package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscan(reader, &N, &M)
	sizes := make([]int, 0, M)
	for i := 0; i < M; i++ {
		sizes = append(sizes, -1)
	}
	for i := 0; i < N; i++ {
		var C, S int
		fmt.Fscan(reader, &C, &S)
		if sizes[C-1] < S {
			sizes[C-1] = S
		}
	}
	for i := 0; i < M; i++ {
		fmt.Print(sizes[i], " ")
	}
}
