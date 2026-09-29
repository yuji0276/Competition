package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	count := 0
	for _, sn := range s {
		if sn == '1' {
			count++
		}
	}
	fmt.Println(count)
}
