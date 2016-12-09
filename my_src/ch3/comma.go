package main

import (
	"fmt"
	"os"
)

func main() {

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + ""
}
