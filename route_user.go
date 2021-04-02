package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	case http.MethodPost:

		var newUser UserData
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(bodyBytes, &newUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println(newUser)
		w.WriteHeader(http.StatusCreated)
	}

}
