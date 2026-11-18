package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	same := true
	seq := true
	for i := range 3 {
		xi := int(x[i] - '0')
		xNext := int(x[i+1] - '0')

		if xi != xNext {
			same = false
		}
		if (xi+1)%10 != xNext {
			seq = false
		}
	}

	if same || seq {
		fmt.Println("Weak")
	} else {
		fmt.Println("Strong")
	}
}
