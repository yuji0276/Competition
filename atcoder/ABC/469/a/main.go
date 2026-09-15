package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int

	fmt.Fscan(reader, &n, &k)
	fmt.Println(n - k + 1)
}
