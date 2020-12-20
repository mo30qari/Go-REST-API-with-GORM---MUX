package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("helloWorld")

}

func createGame(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("createGame: " + mux.Vars(r)["name"] + "@" + mux.Vars(r)["genre"])
	
}

func getAllGames(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("getAllGames")
	
}

func getGame(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("getGame")
	
}

func updateGame(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("updateGame")
	
}

func deleteGame(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("deleteGame")
	
}
