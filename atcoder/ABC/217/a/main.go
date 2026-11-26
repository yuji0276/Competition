package main

import (
	"fmt"
	"slices"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	words := []string{s, t}
	slices.Sort(words)

	if words[0] == s {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
