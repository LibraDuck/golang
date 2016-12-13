package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello"
	str2 := "oloeh"

	if len(str1) != len(str2) {
		fmt.Println("两字符串长度不同")
	}
	for i, v := range ([]byte)(str1) {
		if strings.Count(str1, string(v)) != strings.Count(str2, string(v)) {
			fmt.Printf("[%s]字符的个数不同\n", string(v))
			break
		}
		if i == len(str1)-1 {
			fmt.Println("Yes")
		}
	}
}
