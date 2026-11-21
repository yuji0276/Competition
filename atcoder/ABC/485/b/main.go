package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	drains := make([]string, n)
	for i := range n {
		var drain string
		fmt.Scan(&drain)
		drains[i] = drain
	}

	var faucetStatus bool
	var faucetIdx int
	for i := range n {
		for j := range m {
			if drains[i][j] != 'o' {
				continue
			}
			faucetIdx = i + 1
			faucetStatus = true
		}
	}
	if faucetStatus {
		fmt.Printf("Faucet %d", faucetIdx)
	}
	var drainStatus bool
	var drainIdx int
	for i := range m {
		for j := range n {
			if drains[j][i] == 'o' {
				continue
			}
			drainIdx = i + 1
			drainStatus = true
		}
	}
	if drainStatus {
		fmt.Printf("Drain %d", drainIdx)
	}
}
