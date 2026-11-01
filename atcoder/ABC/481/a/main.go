package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a+b >= 10 {
		fmt.Println("2")
	} else {
		fmt.Println("1")
	}
}
