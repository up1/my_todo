package main

import (
	"net/http"

	"github.com/up1/my_todo/server"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
