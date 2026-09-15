package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dices := make([][]int, 0, 6)
	for i := 0; i < 6; i++ {
		dices[i] = make([]int, 0, 6)
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Fscan(reader, &dices[i][j])
		}
	}
}
