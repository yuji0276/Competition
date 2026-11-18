package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	min := 9
	isNineExist := false
	for range k {
		var a int
		fmt.Scan(&a)
		if a == 9 {
			isNineExist = true
		}
		if min > a {
			min = a
		}
	}
	if isNineExist {
		fmt.Println(min)
	} else {
		fmt.Println(-1)
	}
}
