package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {

	}
	s := os.Args[1]
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:] //记住前开后闭
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	fmt.Println("filename:", s)
}
