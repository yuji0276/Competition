package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	slice := make([]string, n)
	for i := range n {
		var s string
		fmt.Scan(&s)
		slice[i] = s
	}
	target := strings.Repeat("x", m)
	for idx, value := range slice {
		if value == target {
			fmt.Printf("Faucet %v", idx+1)
			return
		}
	}
	for idx := range m {
		if slice[0][idx] == 'x' {
			fmt.Printf("Drain %v", idx+1)
			return
		}
	}
}
