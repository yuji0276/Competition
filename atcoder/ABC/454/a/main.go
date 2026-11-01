package main

import "fmt"

func main() {
	var l, r int

	fmt.Scan(&l, &r)

	count := 0
	for i := l; i <= r; i++ {
		count++
	}
	fmt.Println(count)
}
