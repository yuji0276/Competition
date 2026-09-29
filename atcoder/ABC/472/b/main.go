package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var slice []int
	for i:=0;i<n;i++{
		var a int
		fmt.Scan(&a)
		slice = append(slice, a)
	}

	min_diff := 1000000
	for i:=0;i<n;i++{
		var sum_l int
		var sum_r int
		for j:=0;j<i;j++{
			sum_l += slice[j]
		}
		for j:=i;j<n;j++{
			sum_r += slice[j]
		}
		min_diff = int(math.Min(float64(min_diff), math.Abs(float64(sum_l)-float64(sum_r))))
	}
	fmt.Println(min_diff)
}
