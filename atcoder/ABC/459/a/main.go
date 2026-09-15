package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var X int

	fmt.Fscan(reader, &X)

	S := "HelloWorld"
	result := S[:X-1] + S[X:]
	fmt.Println(result)

}
