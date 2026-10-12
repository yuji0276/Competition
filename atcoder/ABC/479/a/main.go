package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := range n {
		var a int
		fmt.Scan(&a)
		if (i+1)%2 != 0 {
			count += a
		}
	}
	fmt.Println(count)
}
