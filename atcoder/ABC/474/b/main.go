package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	flag := true
	group_number := 0
	for i:=1; i<=n; i++{

		var p int
		fmt.Scan(&p)

		if!((group_number * 10 + 1 <= p)&& (p <= (group_number+1)*10)){
			flag = false
		}
		if (i % 10 == 0){
			group_number++
		}
	}

	if flag{
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
