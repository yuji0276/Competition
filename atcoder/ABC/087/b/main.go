package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				i500 := i * 500
				j100 := j * 100
				k50 := k * 50
				sum := i500 + j100 + k50
				if sum == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
