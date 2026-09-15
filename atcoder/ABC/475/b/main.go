package main

import (
	"fmt"
)

func main(){
	var N int
	ichi := 0
	juu := 0
	hyaku := 0

	fmt.Scan(&N)
	for range(N){
		var a int
		fmt.Scan(&a)

		a = a % 1000

		oturi := 1000 -a
		if (oturi % 1000 == 0){
			continue
		}

		hyaku += oturi / 100
		oturi %= 100

		juu += oturi / 10
		oturi %= 10

		ichi += oturi
	}
	fmt.Println(ichi,juu,hyaku)
}
