package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	var X int

	fmt.Fscan(reader, &N)

	var ans []int
	for i := 0; i < N; i++ {
		var A int

		fmt.Fscan(reader, &A)
		ans = append(ans, A)
	}
	fmt.Fscan(reader, &X)
	fmt.Println(ans[X-1])
}
