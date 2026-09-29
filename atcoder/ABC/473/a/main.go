package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0

	for i:=1;i<=n;i++{
		var a int
		fmt.Scan(&a)
		if i > n/2{
			sum += a
		}
	}
	fmt.Println(sum)
}
