package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Deck struct {
	gorm.Model
	Titre       string
	Description string
	UserID      uint
	Cards       []Card
}

func AllDecks(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	var decks []Deck
	db.Preload("Cards").Find(&decks)
	json.NewEncoder(w).Encode(decks)
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()
	vars := mux.Vars(r)
	titre := vars["titre"]
	description := vars["description"]
	name := vars["name"]
	var user User
	db.Where("name = ? ", name).Find(&user)

	db.Create(&Deck{Titre: titre, Description: description})
	var deck Deck
	db.Where("titre = ? ", titre).Find(&deck)
	deck.UserID = user.ID
	db.Save(&deck)

	fmt.Fprint(w, "NEW DECK created ")
}
func UpdateDeck(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	titre := vars["titre"]
	var deck Deck
	db.Where("titre = ? ", titre).Find(&deck)
	db.Save(&deck)

	fmt.Fprint(w, "deck updated")
}

func DeleteAllDecks(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	var decks []Deck
	db.Find(&decks)

	db.Delete(&decks)
}

//Ajoute une carte dans un deck
func AddCardToDeck(card Card, deck Deck) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	deck.Cards = append(deck.Cards, card)
	db.Save(&deck)

}
