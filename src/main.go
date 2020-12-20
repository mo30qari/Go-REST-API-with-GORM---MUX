package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){

	fmt.Print("gi")
	router := mux.NewRouter()
	router.HandleFunc("/", helloworld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
