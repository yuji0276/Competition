package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	max := 0
	sum := 0
	fmt.Fscan(reader, &N)
	for i := 0; i < N; i++ {
		var L int
		fmt.Fscan(reader, &L)
		if L > max {
			max = L
		}
		sum += L
	}
	if max < sum-max {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
