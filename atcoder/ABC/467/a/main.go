package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var H, W int
	fmt.Fscan(reader, &H, &W)
	if W*10000 >= 25*H*H {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
