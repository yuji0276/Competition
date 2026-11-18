package main

import "fmt"

func main() {
	var n, q int

	fmt.Scan(&n, &q)

	queries := make([]int, q)
	for i := range q {
		var k int
		fmt.Scan(&k)
		queries[i] = k
	}

	var cards []int
	for i := range n {
		cards = append(cards, i+1)
	}

	previousCards := make([]int, n)
	for _, query := range queries {
		copy(previousCards, cards)
		for j := range query {
			cards[j] = 0 - previousCards[query-j-1]
		}
	}
	for _, card := range cards {
		fmt.Println(card)
	}
}
