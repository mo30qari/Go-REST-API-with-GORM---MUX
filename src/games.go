package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

const dbAddress = "root:@/dds_db?charset=utf8&parseTime=True&loc=Local"

type Game struct {
	gorm.Model
	Name string
	Genre string
	Price int
}

func initialMigration(){

	db, err := gorm.Open("mysql", dbAddress)
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

	db, err := gorm.Open("mysql", dbAddress)
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

	db, err := gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open Database!")
	}
	defer db.Close()

	price, _ := strconv.Atoi(mux.Vars(r)["price"])

	game := Game{
		Name: mux.Vars(r)["name"],
		Genre: mux.Vars(r)["genre"],
		Price: price,
	}

	db.Create(&game)

	fmt.Fprintf(w, "The Game sucessfully created!")
	
}

func getGame(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open Database!")
	}
	defer db.Close()

	var game []Game
	db.Where("name = ?", mux.Vars(r)["name"]).Find(&game)

	json.NewEncoder(w).Encode(game)
	
}

func updateGameGenre(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open Database!")
	}
	defer db.Close()

	db.Model(&Game{}).Where("name = ?", mux.Vars(r)["name"]).Update("genre", mux.Vars(r)["genre"])

	fmt.Fprintf(w, "The Game has been updated sucessfully!")
	
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	
	db, err := gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Open DB!")
	}
	defer db.Close()

	db.Where("name = ?", mux.Vars(r)["name"]).Delete(&Game{})

	fmt.Fprintf(w, "The Game has been deleted sucessfully!")
	
}
