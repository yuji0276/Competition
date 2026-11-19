package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	isAC := true
	target := "Hello,World!"

	if s != target {
		isAC = false
	}

	if isAC {
		fmt.Println("AC")
	} else {

		fmt.Println("WA")
	}
}
