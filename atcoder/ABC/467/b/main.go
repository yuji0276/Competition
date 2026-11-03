package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	money := 0
	for range n {
		var a, b int
		var s string
		fmt.Scan(&a, &b, &s)
		if s == "keep" {
			money += b - a
		}
	}
	fmt.Println(money)
}
