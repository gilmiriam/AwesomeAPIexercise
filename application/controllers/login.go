package controllers

import (
	"awesomeProject1/application"
	"encoding/json"
	"net/http"
)

type ResponseToken struct {
	Message     string
	AccessToken string
}

func Login(w http.ResponseWriter, r *http.Request) {
	signedToken := application.GenerateToken()
	w.WriteHeader(http.StatusOK)
	response := ResponseToken{
		Message:     "loginSuccesful",
		AccessToken: signedToken,
	}
	body, _ := json.Marshal(&response)
	w.Write(body)
}
