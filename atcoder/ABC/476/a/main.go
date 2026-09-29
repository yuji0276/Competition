package main

import "fmt"

func main() {
	var s string
  fmt.Scan(&s)

  var t string
	if s[len(s)-1] == 'e' {
		t = s + "r"
	} else {
		t = s + "er"
	}
	fmt.Println(t)
}
