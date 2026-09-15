package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var h, w int
	// h と w を読み込む
	fmt.Fscan(reader, &h, &w)

	// 修正点1: 長さを h に指定して二重スライスを作成
	grid := make([][]byte, h)

	// 修正点2: ループは1つだけでOK
	for i := 0; i < h; i++ {
		var s []byte
		// 1行分の文字列（.#.#. など）をバイト列として読み込む
		fmt.Fscan(reader, &s)
		grid[i] = s
	}
	count := 0
	count += h * w
	fmt.Println(count)
}
