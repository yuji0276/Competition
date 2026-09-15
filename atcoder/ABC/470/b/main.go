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
	slice := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var c int
		fmt.Fscan(reader, &c)
		slice = append(slice, c)
	}
	max := 0
	for i := 0; i < N; i++ {
		count := 0
		for j := 0; j < N; j++ {
			if slice[i] == slice[j] {
				count++
			}
			if max <= count {
				max = count
			}
		}
	}
	fmt.Println(N - max)
}
