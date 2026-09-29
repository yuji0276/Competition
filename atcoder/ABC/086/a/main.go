package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := a * b
	if ans%2 != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
