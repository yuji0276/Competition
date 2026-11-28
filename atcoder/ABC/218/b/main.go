package main

import (
	"fmt"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for range 26 {
		var s int
		fmt.Scan(&s)
		fmt.Print(string(alphabet[s-1]))
	}
}
