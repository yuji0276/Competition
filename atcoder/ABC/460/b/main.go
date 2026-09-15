package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkSharePoint(X1 int, Y1 int, R1 int, X2 int, Y2 int, R2 int) {
	dx := X2 - X1
	dy := Y2 - Y1
	dd := dx*dx + dy*dy
	if (dd <= (R1+R2)*(R1+R2)) && ((R1-R2)*(R1-R2) <= dd) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		var X1, Y1, R1, X2, Y2, R2 int
		fmt.Fscan(reader, &X1, &Y1, &R1, &X2, &Y2, &R2)
		checkSharePoint(X1, Y1, R1, X2, Y2, R2)
	}
}
