package main

import "fmt"

func main() {

	contests := make(map[string]struct{})

	contests["ABC"] = struct{}{}
	contests["ARC"] = struct{}{}
	contests["AGC"] = struct{}{}
	contests["AHC"] = struct{}{}

	for range 3 {
		var s string
		fmt.Scan(&s)
		delete(contests, s)
	}

	for contest := range contests {
		fmt.Println(contest)
	}
}
