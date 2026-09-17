package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x{
	case 1:
		fmt.Println(2)
	case 2:
		fmt.Println(3)
	case 3:
		fmt.Println(1)
	}
}
