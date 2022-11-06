package main

import (
	"awesomeProject1/application/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	routerParent := mux.NewRouter().StrictSlash(true)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	routerParent.HandleFunc("/ping", controllers.Ping).Methods("GET")
	routerParent.HandleFunc("/login", controllers.Login).Methods("POST")
	routerParent.HandleFunc("/uploadFile", controllers.UploadFile).Methods("POST")
	//routerParent.HandleFunc("/user", controllers.Login).Methods("GET")
	//routerParent.HandleFunc("/refresh", controllers.Login).Methods("POST")

	handler := corsMiddleware.Handler(routerParent)

	http.ListenAndServe(":8080", handler)
}
