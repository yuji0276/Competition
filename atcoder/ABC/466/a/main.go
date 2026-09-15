package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	isUnhappy := true
	fmt.Fscan(reader, &N)
	for i := 0; i < N; i++ {
		var X int
		fmt.Fscan(reader, &X)
		if X >= 0 {
			isUnhappy = false
			break
		}
	}
	if isUnhappy {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
