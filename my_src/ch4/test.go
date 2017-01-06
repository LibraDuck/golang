package main

import (
	"fmt"
)

func main() {
	var arry [13]int
	Q2 := arry[4:7]
	fmt.Printf(" len = %d \n cap = %d \n", len(Q2), cap(Q2))
}
