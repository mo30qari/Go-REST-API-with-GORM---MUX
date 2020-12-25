package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB
var err error
const dbAddress = "root:@/dds_db?charset=utf8&parseTime=True&loc=Local"

type Game struct {
	gorm.Model
	Name string
	Genre string
	Price int
}

func initialMigration(){

	db, err = gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open Database!")
	}
	defer db.Close()
	
	db.AutoMigrate(&Game{})
}

func welcome(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("This is the first page of the API. You can't see anything here...")
	
}

func getAllGames(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getAllGames")

	db, err = gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open Database!")
	}
	defer db.Close()

	var games []Game
	db.Find(&games)

	json.NewEncoder(w).Encode(games)

}

func createGame(w http.ResponseWriter, r *http.Request) {

	fmt.Println(1)
	fmt.Println("createGame: " + mux.Vars(r)["name"] + "@" + mux.Vars(r)["genre"])
	
}



func getGame(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("getGame")
	
}

func updateGame(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("updateGame")
	
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("deleteGame")
	
}
