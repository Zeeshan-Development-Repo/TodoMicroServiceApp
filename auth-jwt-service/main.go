package main

import (
	"auth-jwt-service/login"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var secretKey = []byte("secret-key")

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", login.ProtectedHandler).Methods("GET")

	fmt.Println("Starting the server")
	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
	fmt.Println("Server Started at :4000")
}
