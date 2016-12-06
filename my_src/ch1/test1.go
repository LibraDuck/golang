package main

import (
	"fmt"
)

func main() {
	//num := 10
	as := make(map[string]interface{})
	as["tag"] = "good"
	fmt.Printf("%T\n", as["as"])
}
