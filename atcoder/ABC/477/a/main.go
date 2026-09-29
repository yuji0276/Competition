package main

import (
	"fmt"
)

func main() {
	var c string
	fmt.Scan(&c)
	switch c {
	case "B":
		fmt.Println("Y")
	case "Y":
		fmt.Println("R")
	case "R":
		fmt.Println("B")
	}
}
