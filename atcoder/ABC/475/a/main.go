package main

import (
	"fmt"
)

func main(){
	var s string
	result := ""
	fmt.Scan(&s)
	for i,char := range(s){
		if i == len(s)-1{
			result += string(char)
		}else{
			result += string(char)+"o"
		}
	}
	fmt.Println(result)
}
