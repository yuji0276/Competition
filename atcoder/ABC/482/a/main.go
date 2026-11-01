package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)
	count := 0
	for _, value := range s {
		if value == 'o' {
			count++
		}
	}
	if count%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
