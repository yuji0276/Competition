package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	firstNames := make([]string, n)
	lastNames := make([]string, n)
	for i := range n {
		var firstName string
		var lastName string
		fmt.Scan(&lastName, &firstName)
		firstNames[i] = firstName
		lastNames[i] = lastName
	}

	flag := false
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (firstNames[i] == firstNames[j]) && (lastNames[i] == lastNames[j]) {
				flag = true
				break
			}
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
