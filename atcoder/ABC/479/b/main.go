package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	renzoku := 0
	saikoukiroku := 0
	for i := 0; i < n; i++ {
		if s[i] == 'S' {
			renzoku = 0
		} else {
			renzoku++
		}
		if renzoku > saikoukiroku {
			saikoukiroku = renzoku
		}
	}
	fmt.Println(saikoukiroku)
}
