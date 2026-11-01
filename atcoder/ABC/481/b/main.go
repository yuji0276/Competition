package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	points := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		points[i] = a
	}

	count := 0
	skipedAt := 101
	for i := 0; i < n-1; i++ {
		if i == skipedAt {
			continue
		}
		switch k {
		case points[i]:
			skipedAt = i
			count++
			continue
		case points[i+1]:
			skipedAt = i + 1
			count++
			continue
		}
		if (min(points[i], points[i+1]) <= k) && (k <= max(points[i], points[i+1])) {
			count++
		}
	}
	fmt.Println(count)
}
