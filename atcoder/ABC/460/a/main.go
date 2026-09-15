package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N, M int

	fmt.Fscan(reader, &N, &M)

	count := 0
	for {
		x := N % M
		M = x
		count += 1
		if M == 0 {
			break
		}
	}
	fmt.Println(count)
}
