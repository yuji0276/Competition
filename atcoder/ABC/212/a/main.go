package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if 0 < a && b == 0 {
		fmt.Println("Gold")
	}
	if a == 0 && b > 0 {
		fmt.Println("Silver")
	}
	if 0 < a && 0 < b {
		fmt.Println("Alloy")
	}
}
