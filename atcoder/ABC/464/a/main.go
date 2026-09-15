package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var S string
	var count int
	var isEastWin bool

	fmt.Fscan(reader, &S)
	for i := range S {
		if S[i] == 'E' {
			count += 1
		}
	}

	n := utf8.RuneCountInString(S)
	if count > n/2 {
		isEastWin = true
	}
	if isEastWin {
		fmt.Println("East")
	} else {
		fmt.Println("West")
	}
}
