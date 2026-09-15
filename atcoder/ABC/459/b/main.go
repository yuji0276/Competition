package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(reader, &N)
	keypad := "22233344455566677778889999"

	ans := ""
	for i := 0; i < N; i++ {
		var S string
		fmt.Fscan(reader, &S)
		ans += string(keypad[S[0]-'a'])
	}
	fmt.Println(ans)
}
