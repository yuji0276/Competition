package main

import (
	"bufio"
	"fmt"
	"os"
)

func NumCheck(str string) bool {
	for _, r := range str {
		if '0' <= r && r <= '9' {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var S string
	var ans string

	fmt.Fscan(reader, &S)

	for i := range S {
		if NumCheck(string(S[i])) {
			ans = ans + string(S[i])
		}
	}
	fmt.Println(ans)
}
