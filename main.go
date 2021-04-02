package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ok")
	handleRequest()
}

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success")
}

func handleRequest() {
	http.HandleFunc("/", users)
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8080", nil)
}
