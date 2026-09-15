package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(reader, &N)
	for i := 1; i <= N; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}
