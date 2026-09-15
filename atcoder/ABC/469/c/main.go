package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	var s string

	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &s)

	xPos := make([]int, 0, n)
	for i := range n {
		if s[i] == 'x' {
			xPos = append(xPos, i+1)
		}
	}

	buf := make([]byte, 0, 8)
	for k := 1; k <= n; k++ {
		ans := n
		if k <= len(xPos) {
			ans = xPos[k-1]
		}

		buf = strconv.AppendInt(buf[:0], int64(ans), 10)
		buf = append(buf, '\n')
		writer.Write(buf)
	}
}
