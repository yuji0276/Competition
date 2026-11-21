package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if ((a > 0) && (b < 0)) || ((a < 0) && (b > 0)) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
