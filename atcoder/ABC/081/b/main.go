package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var numbers []int
	for range n {
		var a int
		fmt.Scan(&a)
		numbers = append(numbers, a)
	}
	count := 0
	flag := true
	for flag {
		for i, each := range numbers {
			if each%2 == 0 {
				numbers[i] = each / 2
			} else {
				flag = false
				count--
				break
			}
		}
		count++
	}
	fmt.Println(count)
}
