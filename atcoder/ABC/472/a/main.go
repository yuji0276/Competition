package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	r := []rune(s)
	for i, t := range(s){
		if (t != 'A'){
			r[i] = '.'
		}
	}
	s = string(r)
	fmt.Println(s)
}
