package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ok")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	userResponse := UserData{
		Userid:   28566777,
		Email:    "training@nextflow.in.th",
		Password: "1234",
	}
	json.NewEncoder(w).Encode(userResponse)
}
