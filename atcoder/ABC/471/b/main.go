package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	results := make([]string, n)

	for i := range n {
		var result string
		fmt.Scan(&result)
		result = strings.ToLower(result)
		results[i] = result
	}

	max := 0
	for i := range n {
		count := 0
		for j := range n {
			if results[i] == results[j] {
				count++
			}
		}
		if count > max {
			max = count
		}
	}

	fmt.Println(max)
}
