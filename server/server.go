package main

import (
	"fmt"
	"net/http"
)

type myhttp int

func (m myhttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go")
}

func main() {
	var d myhttp
	http.ListenAndServe(":8080", d)
}
