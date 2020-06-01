package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Card struct {
	gorm.Model
	Recto  string
	Verso  string
	DeckID uint
}

func AllCards(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	var cards []Card
	db.Find(&cards)
	json.NewEncoder(w).Encode(cards)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	recto := vars["recto"]
	verso := vars["verso"]
	titre := vars["titre"]
	var deck Deck

	db.Where("titre = ? ", titre).Find(&deck)
	db.Create(&Card{Recto: recto, Verso: verso})
	var card Card
	db.Where("recto = ? ", recto).Find(&card)

	card.DeckID = deck.ID
	db.Preload("Cards").Find(&deck)
	db.Save(&card)
	AddCardToDeck(card, deck)

	fmt.Fprint(w, "NEW CArd created ")
}
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	recto := vars["recto"]
	verso := vars["verso"]

	var card Card
	db.Where("recto = ? ", recto).Find(&card)

	card.Verso = verso

	db.Save(&card)

	fmt.Fprint(w, "card updated")
}

func DeleteAllCards(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	var cards []Card
	db.Find(&cards)

	db.Delete(&cards)
}

//Renvoi toutes les cards d'un deck
func GetCardFromDeck(deck Deck) []Card {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()
	var cards []Card
	db.Preload("Cards").Find(&deck)
	return cards

}
