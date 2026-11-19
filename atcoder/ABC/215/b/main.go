package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	k := 0
	flag := true
	for flag {
		if int(math.Pow(2, float64(k))) > n {
			k--
			flag = false
			break
		} else {
			k++
		}
	}
	fmt.Println(k)
}
