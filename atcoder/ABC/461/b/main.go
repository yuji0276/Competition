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

	ax := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(reader, &ax[i])
	}

	kikori := make([]int, N+1)
	fmt.Fscan(reader, &kikori)
	for i := 1; i <= N; i++ {
		fmt.Fscan(reader, &kikori[i])
	}

	isHonest := true
	for i := 1; i <= N; i++ {
		if i != kikori[ax[i]] {
			isHonest = false
			break
		}
	}
	if isHonest {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
