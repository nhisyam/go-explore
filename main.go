package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome, "+r.URL.Path[1:]+"!")
}

func main() {
	fmt.Println("Hello world!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
