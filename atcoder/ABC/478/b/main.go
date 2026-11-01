package main

import "fmt"

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	values := make([]int, n)
	for i := range n {
		fmt.Scan(&values[i])
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				valueSum := values[i] + values[j] + values[k]
				priceSum := i + j + k + 3
				fmt.Println(i, j, k)
				if (priceSum <= v) && (valueSum >= max) {
					max = valueSum
				}
			}
		}
	}
	fmt.Println(max)
}
