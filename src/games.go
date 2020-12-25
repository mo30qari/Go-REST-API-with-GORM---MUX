package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB
var err error

type Game struct {
	gorm.Model
	Name string
	Genre string
	Price int
}

func initialMigration(){
	
	db, err = gorm.Open("mysql", "root:@/dds_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open Database!")
	}
	defer db.Close()
	
	db.AutoMigrate(&Game{})
}

func welcome(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("helloWorld")
	
}

func createGame(w http.ResponseWriter, r *http.Request) {

	fmt.Println(1)
	fmt.Println("createGame: " + mux.Vars(r)["name"] + "@" + mux.Vars(r)["genre"])
	
}

func getAllGames(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("getAllGames")
	
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
