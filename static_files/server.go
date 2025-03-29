package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server listening on port :8080")

	http.ListenAndServe(":8080", nil)
}
