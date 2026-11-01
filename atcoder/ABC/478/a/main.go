package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	slice := make([]int, n)

	for i := range m {
		slice[i%n]++
	}
	for i := range n {
		fmt.Println(slice[i])
	}
}
