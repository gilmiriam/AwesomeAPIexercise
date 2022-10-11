package main

import (
	"awesomeProject1/application/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	routerParent := mux.NewRouter().StrictSlash(true)

	routerParent.HandleFunc("/ping", controllers.Ping).Methods("GET")

	http.ListenAndServe(":8080", routerParent)
}
