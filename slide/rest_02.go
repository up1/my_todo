package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", addHanler)
	http.ListenAndServe(":8080", nil)
}

func addHanler(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(w, request)
		return
	}
	fmt.Fprintf(w, "Welcome to the home page!")
}
