package main

import (
	"fmt"
	"log"
	"model"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "liste des decks ")
}

func main() {

	model.InitialMigration()
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
