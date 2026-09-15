package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var X int
	fmt.Fscan(reader, &X)
	var isSum bool
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				if i+j+k == X {
					isSum = true
					break
				}
			}
		}
	}
	if isSum {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
