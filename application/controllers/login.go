package controllers

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("{token: test123}")
	w.Write(body)
}
