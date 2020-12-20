package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func main(){

	handleRequests()

}

func handleRequests(){
	
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/game/{name}/{genre}", createGame).Methods("GET")
	router.HandleFunc("/games", getAllGames).Methods("GET")
	router.HandleFunc("/game/{name}", getGame).Methods("GET")
	router.HandleFunc("/game/{name}", updateGame).Methods("PUT")
	router.HandleFunc("/game/{name}", deleteGame).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
	
}
