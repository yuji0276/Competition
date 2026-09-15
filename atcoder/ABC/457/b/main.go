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

	arrays := make([][]int, N)
	for i := 0; i < N; i++ {
		var L int
		fmt.Fscan(reader, &L)
		arrays[i] = make([]int, L)
		for j := 0; j < L; j++ {
			var A int
			fmt.Fscan(reader, &A)
			arrays[i][j] = A
		}
	}
	var X, Y int
	fmt.Fscan(reader, &X, &Y)
	fmt.Println(arrays[X-1][Y-1])
}
