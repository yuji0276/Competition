package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var X string

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &N, &X)

	idx := X[0] - 'A'
	var isVacant bool
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(reader, &s)
		if s[idx] == 'o' {
			isVacant = true
		}
	}
	if isVacant {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
