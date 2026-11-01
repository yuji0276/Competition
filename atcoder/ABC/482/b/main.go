package main

import (
	"fmt"
)

func main() {
	bingo := make([]string, 5)
	for i := range bingo {
		var s string
		fmt.Scan(&s)
		bingo[i] = s
	}

	min := 5

	for i := range bingo {
		var row string
		counts := 0
		for j := range bingo {
			row = bingo[i]
			if row[j] == '.' {
				counts++
			}
		}
		if counts < min {
			min = counts
		}
	}

	for i := range bingo {
		counts := 0
		var row string
		for j := range bingo {
			row = bingo[j]
			if row[i] == '.' {
				counts++
			}
		}
		if counts < min {
			min = counts
		}
	}

	count := 0
	for i := range bingo {
		row := bingo[i]
		if row[i] == '.' {
			count++
		}
	}
	if count < min {
		min = count
	}

	count = 0
	for i := range bingo {
		row := bingo[i]
		if row[4-i] == '.' {
			count++
		}
	}
	if count < min {
		min = count
	}

	fmt.Println(min)
}
