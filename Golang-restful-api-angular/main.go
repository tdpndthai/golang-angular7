package main

import (
	"fmt"
	"golang-restful-api-angular/api"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//fmt.Println("hello")
	router := mux.NewRouter()
	router.HandleFunc("/api/product/findall", api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/find/{id}", api.FindById).Methods("GET")
	router.HandleFunc("/api/product/create", api.Create).Methods("POST")
	router.HandleFunc("/api/product/delete/{id}", api.Delete).Methods("DELETE")
	router.HandleFunc("/api/product/update", api.Update).Methods("PUT")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "User-Agent", "Origin", "auth_token", "X-CSRF-Token"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})
	err := http.ListenAndServe(":6535", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Listening 6535.....")
	}
}
