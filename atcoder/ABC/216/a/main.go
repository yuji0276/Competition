package main

import "fmt"

func main() {
	var fig float64
	fmt.Scan(&fig)
	x := int(fig)
	y := int((fig - float64(x)) * 10)
	if (0 <= y) && (y <= 2) {
		fmt.Printf("%v-", x)
	} else if (3 <= y) && (y <= 6) {
		fmt.Printf("%v", x)
	} else {
		fmt.Printf("%v+", x)
	}
}
