package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)

	names := make(map[string]int, q)
	for i := 0; i < q; i++ {
		var name string
		fmt.Scan(&name)
		names[name] = 0
	}

	count := 1
	for i := 0; i < q; i++ {
		for j := 0; j < q; j++ {
			if names[i] == names[j] && i != j {
				fmt.Println()
			} else {

			}
		}
	}
}
