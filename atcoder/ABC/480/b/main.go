package main

import "fmt"

func isValid(candicate, vote string) bool {
	r1, r2 := []rune(candicate), []rune(vote)
	if len(r1) != len(r2) {
		return false
	}

	diffCount := 0
	for i := range len(r1) {
		if r1[i] != r2[i] {
			diffCount++
		}
	}
	if diffCount == 1 || diffCount == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	candicates := make([]string, n)
	votes := make([]string, m)

	for i := range n {
		var candicate string
		fmt.Scan(&candicate)
		candicates[i] = candicate
	}

	for i := range m {
		var vote string
		fmt.Scan(&vote)
		votes[i] = vote
	}
	counts := make([]int, n)
	for j := 0; j < m; j++ {
		exact := -1
		for i := 0; i < n; i++ {
			if candicates[i] == votes[j] {
				exact = i
				break
			}
		}
		if exact >= 0 {
			counts[exact]++
			continue
		}

		nearCount := 0
		nearIndex := -1
		for i := 0; i < n; i++ {
			if isValid(candicates[i], votes[j]) {
				nearCount++
				nearIndex = i
			}
		}
		if nearCount == 1 {
			counts[nearIndex]++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(counts[i])
	}
}
