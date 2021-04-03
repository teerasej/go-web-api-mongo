package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	var mongoServer = "localhost"
	var envMongoServer = os.Getenv("MONGO_SERVER")

	if envMongoServer != "" {
		mongoServer = envMongoServer
	}

	err := mgm.SetDefaultConfig(nil, "nextflow", options.Client().ApplyURI("mongodb://"+mongoServer+":27017"))
	if err != nil {
		panic(err)
	}
}

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
			w.Write([]byte("cannot read body"))
		}
		err = json.Unmarshal(bodyBytes, &newUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("cannot parse JSON"))
		}

		err = mgm.Coll(&newUser).Create(&newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("cannot create new user"))
		}

		fmt.Println(newUser)

		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message":"success"}`))
	}

}
