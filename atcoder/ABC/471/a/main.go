package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := a + b
	diff := a - b
	pro := a * b

	flag := false
	if (sum == 9) || (diff == 9) || (pro == 9) {
		flag = true
	}
	if b != 0 {
		if a%b == 0 && a/b == 9 {
			flag = true
		}
	}
	if flag {
		fmt.Println("Nine")
	} else {
		fmt.Println("Nein")
	}
}
