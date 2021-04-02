package main

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userResponse := UserData{
			Userid:   28566777,
			Email:    "training@nextflow.in.th",
			Password: "1234",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userResponse)
	}

}
