package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var M int
	var D int
	var S string
	fmt.Fscan(reader, &M, &D)
	fmt.Fscan(reader, &S)

	diff := make([]int, M+1)
	for i := 0; i < M; i++ {
		if S[i] != 'G' {
			continue
		}
		l := max(i-D, 0)
		r := min(i+D, M-1)
		diff[l]++
		diff[r+1]--
	}

	count := 0
	cur := 0
	for i := 0; i < M; i++ {
		cur = cur + diff[i]
		if cur == 0 {
			count = count + 1
		}
	}
	fmt.Println(count)
}
