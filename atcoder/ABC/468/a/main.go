package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	var A []int
	fmt.Fscan(reader, &N)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(reader, &a)
		A = append(A, a)
	}
	count := 0
	for i := 0; i < N-2; i++ {
		if (A[i] < A[i+1]) && (A[i+1] > A[i+2]) {
			count = count + 1
		}
	}
	fmt.Println(count)
}
