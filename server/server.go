package main

import (
	"fmt"
	"net/http"
)

type myhttp int

func c(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello c")
}

func d(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello d")
}

func main() {
	http.HandleFunc("/c", c)
	http.HandleFunc("/d", d)
	http.ListenAndServe(":8080", nil)
}
