package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c int

	fmt.Fscan(reader, &a, &b, &c)
	if (a != b) && (b == c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
