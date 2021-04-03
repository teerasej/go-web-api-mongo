package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is starting")
	fmt.Println("Server is running at port 8080")
	handleRequest()
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func handleRequest() {
	http.HandleFunc("/", users)
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8080", nil)
}
