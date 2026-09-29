package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0

	var cards []int

	for range(n){
		var a int
		fmt.Scan(&a)
		cards = append(cards, a)
	}

	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if (cards[i] == cards[j]) && (i != j){
				cards[i] = 0
				cards[j] = 0
			}
		}
		sum += cards[i]
	}
	fmt.Println(sum)
}
