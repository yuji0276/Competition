package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var S string
	var N int
	fmt.Fscan(reader, &S)
	fmt.Fscan(reader, &N)

	ans := string(S[N : len(S)-N])
	fmt.Println(ans)
}
