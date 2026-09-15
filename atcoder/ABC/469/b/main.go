package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	s := make([]byte, 0, n)

	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &s)
	count := 0
	for i := range n {
		isVacant := false
		flagLeft := false
		flagRight := false

		if s[i] == 'x' {
			isVacant = true
		}

		if i == 0 {
			flagLeft = true
		} else {
			if s[i-1] == 'x' {
				flagLeft = true
			}
		}

		if i == n-1 {
			flagRight = true
		} else {
			if s[i+1] == 'x' {
				flagRight = true
			}
		}

		if (flagLeft && isVacant) && flagRight {
			count++
		}
	}
	fmt.Println(count)
}
