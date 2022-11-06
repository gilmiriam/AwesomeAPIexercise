package controllers

import (
	"encoding/json"
	"net/http"
)

type UserCredentials struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(user)
	w.Write(body)
}
