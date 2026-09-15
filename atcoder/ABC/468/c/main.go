package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int

	fmt.Fscan(reader, &N)
	P := make([]int, 0, N)
	Q := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var p int
		fmt.Fscan(reader, &p)
		P = append(P, p)
	}
	for i := 0; i < N; i++ {
		var q int
		fmt.Fscan(reader, &q)
		Q = append(Q, q)
	}
	count := 0
	for i := 0; i < N; i++ {
	}
}
