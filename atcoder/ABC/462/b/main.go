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

	recv := make([][]int, N+1)

	for i := 1; i <= N; i++ {
		var K int
		fmt.Fscan(reader, &K)
		for j := 0; j < K; j++ {
			var a int
			fmt.Fscan(reader, &a)
			recv[a] = append(recv[a], i)
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := 1; i <= N; i++ {
		fmt.Fprint(writer, len(recv[i]))
		for j := range len(recv[i]) {
			fmt.Fprint(writer, " ", recv[i][j])
		}
		fmt.Fprintln(writer)
	}
}
