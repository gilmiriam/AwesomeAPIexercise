package controllers

import (
	"encoding/json"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	f, err := os.OpenFile("./etc/files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("{message: File" + fileName + " Uploaded successfully}")
	w.Write(body)
}
