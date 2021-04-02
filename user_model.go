package main

import (
	"github.com/kamva/mgm/v3"
)

type UserData struct {
	mgm.DefaultModel `bson:",inline"`
	Userid           int    `json:"id"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}
