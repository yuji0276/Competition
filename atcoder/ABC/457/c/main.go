package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, K int
	fmt.Fscan(reader, &N, &K)
	K--

	arrayA := make([][]int, N)
	for i := 0; i < N; i++ {
		var L int
		fmt.Fscan(reader, &L)
		arrayA[i] = make([]int, L)
		for j := 0; j < L; j++ {
			var A int
			fmt.Fscan(reader, &A)
			arrayA[i][j] = A
		}
	}
	arrayC := make([]int, N)
	for i := 0; i < N; i++ {
		var C int
		fmt.Fscan(reader, &C)
		arrayC[i] = C
	}

	for i := 0; i < N; i++ {
		L := len(arrayA[i])
		if K < arrayC[i]*L {
			fmt.Fprintln(writer, arrayA[i][K%L])
			break
		}
		K -= arrayC[i] * L
	}
}
