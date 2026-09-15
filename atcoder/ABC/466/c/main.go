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

	var idx int
	var ans string

	for i := 1; i <= N; i++ {
		fmt.Printf("? %d %d\n", i, N-i+1)

		fmt.Fscan(reader, &ans)
		if ans == "Yes" {
			idx = i
			break
		}
	}
	fmt.Printf("? %d %d\n", idx+2, N-idx+1)
	fmt.Fscan(reader, &ans)
}
