package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	scores := make([]int, n)

	for i := range n {
		var a int
		fmt.Scan(&a)
		scores[i] = a
	}

	last := 0
	for i := range n {
		if last < scores[i] {
			last = scores[i]
		}
	}
	booby := 0
	for i := range n {
		if (booby < scores[i]) && (scores[i] < last) {
			booby = scores[i]
		}
	}

	for i := range n {
		if booby == scores[i] {
			fmt.Println(i + 1)
		}
	}
}
