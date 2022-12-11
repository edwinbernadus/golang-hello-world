package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write a response with a "Hello World" message
		fmt.Fprintf(w, "Hello World")
	})

	// Start the server on port 3000
	http.ListenAndServe(":3000", nil)
}
