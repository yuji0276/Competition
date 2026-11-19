package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	count := 0
	for i := 0; i < 8; i++ {
		digitA := a % 2
		digitB := b % 2
		a /= 2
		b /= 2
		if digitA != digitB {
			count += int(math.Pow(2, float64(i)))
		}
	}
	fmt.Println(count)
}
