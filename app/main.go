package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("MESSAGE")
	fmt.Fprintf(w, message+"\nThis is version 3\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":"+port, nil)
}
