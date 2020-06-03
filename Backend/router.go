package main

import (
	"model"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", helloworld).Methods("GET")

	//Route User
	myRouter.HandleFunc("/users", model.AllUsers).Methods("GET")

	myRouter.HandleFunc("/user/{name}", model.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", model.UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/user/{name}/{email}/{password}", model.CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/GetUser/{email}", model.GetUser).Methods("GET")
	myRouter.HandleFunc("/user/connection/{email}/{password}", model.ConnectionUser).Methods("GET")

	//Route Deck
	myRouter.HandleFunc("/decks", model.AllDecks).Methods("GET")
	myRouter.HandleFunc("/deck/{titre}/{description}/{name}", model.CreateDeck).Methods("POST")
	myRouter.HandleFunc("/deck/{id}/{titre}", model.UpdateDeck).Methods("PUT")
	myRouter.HandleFunc("/deck/deleteall", model.DeleteAllDecks).Methods("DELETE")

	//Route Card
	myRouter.HandleFunc("/cards", model.AllCards).Methods("GET")
	myRouter.HandleFunc("/card/{recto}/{verso}/{titre}", model.CreateCard).Methods("POST")
	myRouter.HandleFunc("/card/{titre}/{verso}", model.UpdateCard).Methods("PUT")
	myRouter.HandleFunc("/card/deleteall", model.DeleteAllCards).Methods("DELETE")

	return myRouter

}
