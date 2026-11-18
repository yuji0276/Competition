package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	primes := make([]int, m)
	secondaries := make([]int, m)

	for i := range m {
		var a, b int
		fmt.Scan(&a, &b)
		primes[i] = a
		secondaries[i] = b
	}

	for i := range n {
		votes := make([]int, n)
		for j := range m {
			if primes[j] == i+1 {
				votes[secondaries[j]]++
			} else {
				votes[primes[j]]++
			}
		}
		fmt.Println("votes:", votes)
		max := 0
		for k := range len(votes) {
			if max < votes[k] {
				max = k
			}
		}
		fmt.Println(max)
	}

}
