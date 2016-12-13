package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] //加强对递归的了解，可以让代码更好得自动化
}
