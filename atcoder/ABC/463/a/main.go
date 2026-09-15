package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y int

	fmt.Fscan(reader, &x, &y)
	if 9*x == 16*y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
