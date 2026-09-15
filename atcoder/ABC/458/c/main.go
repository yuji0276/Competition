package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var S string

	fmt.Fscan(reader, &S)
	count := 0
	for j := 0; j < len(S); j++ {
		if S[j] == 'C' {
			count += min(j+1, len(S)-j)
		}
	}
	fmt.Println(count)
}
