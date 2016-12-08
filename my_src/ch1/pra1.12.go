package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	fmt.Println("cycles begin:", cycles)

	cycles, _ = strconv.Atoi(r.FormValue("cycles"))
	fmt.Println("cycles end:", cycles)
}
