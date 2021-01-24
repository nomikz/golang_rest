package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("someone came in again")
		_, _ = fmt.Fprintf(w, "<h1>Hello kitty</h1>")
	})
	_ = http.ListenAndServe(":3000", nil)
}