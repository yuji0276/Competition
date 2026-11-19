package main

import "fmt"

func main() {
	var s, t int
	fmt.Scan(&s, &t)

	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			for k := 0; k < 1000; k++ {
				if (i+j+k <= s) && (i*j*k <= t) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
