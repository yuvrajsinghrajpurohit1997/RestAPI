package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/new", NewUser).Methods("POST")
	myRouter.HandleFunc("/users", GetUser).Methods("GET")
	myRouter.HandleFunc("/users/{id}", GetUserID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
func main() {
	fmt.Println("Go ORM Practice")
	handlerequest()

}
