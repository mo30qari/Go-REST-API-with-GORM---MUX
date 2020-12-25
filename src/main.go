package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func main() {
	
	initialMigration()
	
	handleRequests()
	
}

func handleRequests() {
	
	router := mux.NewRouter()
	
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/game/{name}/{genre}/{price}", createGame).Methods("GET")
	router.HandleFunc("/games", getAllGames).Methods("GET")
	router.HandleFunc("/game/{name}", getGame).Methods("GET")
	router.HandleFunc("/game/{name}", updateGame).Methods("PUT")
	router.HandleFunc("/game/{name}", deleteGame).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", router))
	
}
