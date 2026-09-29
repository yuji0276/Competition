package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, d int
	fmt.Fscan(in, &n, &d)

	people := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &people[i])
	}

	var standingOutliers []int
	for i := 0; i < n; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if i != j && abs(people[i]-people[j]) < d {
				ok = false
				break
			}
		}
		if ok {
			standingOutliers = append(standingOutliers, i+1)
		}
	}

	fmt.Fprintln(out, len(standingOutliers))
	for _, e := range standingOutliers {
		fmt.Fprint(out, e, " ")
	}
	fmt.Fprintln(out)
}
