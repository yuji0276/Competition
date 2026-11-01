package main

import "fmt"

func main() {
	var s1 string
	var s2 string

	fmt.Scan(&s1, &s2)

	if ((s1[0] == 'A') && (s2[1] == 'D')) || ((s1[1] == 'A') && (s2[0] == 'D')) || ((s1[0] == 'D') && (s2[1] == 'A')) || ((s1[1] == 'D') && (s2[0] == 'A')) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
