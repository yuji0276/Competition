package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var A, D int
	fmt.Fscan(reader, &A, &D)

	var iSDefenced bool

	if D >= A {
		iSDefenced = true
	}

	if iSDefenced {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
