package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var A, B int
	fmt.Fscan(reader, &A, &B)

	if A > B*2/3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
