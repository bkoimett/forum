package main

import (
	"fmt"
	"net/http"

	
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, the server appears to be on"))
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}