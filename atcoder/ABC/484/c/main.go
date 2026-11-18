package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)
	readInt := func() int {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		return v
	}

	n := readInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		a[i] = readInt()
	}
	for i := range n {
		b[i] = readInt()
	}

	l := 0
	for l < n && a[l] == b[l] {
		l++
	}
	if l == n {
		fmt.Println("Yes")
		return
	}
	r := n - 1
	for a[r] == b[r] {
		r--
	}

	ok1 := a[l] == b[r]
	for k := 1 + l; ok1 && k < r; k++ {
		if b[k] != a[k+1] {
			ok1 = false
		}
	}
	ok2 := a[r] == b[l]
	for k := 1 + l; ok2 && k <= r; k++ {
		if b[k] != a[k-1] {
			ok2 = false
		}
	}

	if ok1 || ok2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
