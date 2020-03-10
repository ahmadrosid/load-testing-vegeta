package main

import (
	"fmt"
	"net/http"
)

var count = 0

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Printf("%d\n", count)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
