package main

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userResponse := UserData{
		Userid:   28566777,
		Email:    "training@nextflow.in.th",
		Password: "1234",
	}
	json.NewEncoder(w).Encode(userResponse)
}
