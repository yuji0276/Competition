package main

import "fmt"

func main() {
	var n int
	var s string
	var t string
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Scan(&t)
	flag := true
	for i, char := range(t){
		if (char == '*'){
			continue
		}else{
			if (char != rune(s[i])){
				flag = false
			}
		}
	}
		if flag{
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
}
