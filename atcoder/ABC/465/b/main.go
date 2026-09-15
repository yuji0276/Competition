package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var X, Y, L, R, A, B int
	var cost int
	fmt.Fscan(reader, &X, &Y, &L, &R, &A, &B)
	if (B <= L) || (R <= A) {
		cost = (B - A) * Y
	} else if (A <= L) && (R <= B) {
		cost = (L-A)*Y + (R-L)*X + (B-R)*Y
	} else if (A <= L) && (L <= B) {
		cost = (L-A)*Y + (B-L)*X
	} else if (L <= A) && (B <= R) {
		cost = (B - A) * X
	} else if (A <= R) && (R <= B) {
		cost = (R-A)*X + (B-R)*Y
	}
	fmt.Println(cost)
}
