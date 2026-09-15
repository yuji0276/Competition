package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var X, Y, L, R, A, B int
	var cost int
	fmt.Fscan(reader, &X, &Y, &L, &R, &A, &B)

	for i := A; i < B; i++ {
		if (L <= i) && (i <= R-1) {
			cost += X
		} else {
			cost += Y
		}
	}
	fmt.Println(cost)
}
