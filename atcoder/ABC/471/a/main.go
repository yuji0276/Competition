package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var a, b int

	fmt.Fscan(r, &a, &b)

	if a+b == 9 {
		fmt.Println("Nine")
	} else if a-b == 9 {
		fmt.Println("Nine")
	} else if a*b == 9 {
		fmt.Println("Nine")
	} else if a == 9*b {
		fmt.Println("Nine")
	} else {
		fmt.Println("Nein")
	}
}
