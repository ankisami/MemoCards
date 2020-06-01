package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ConnectionString = "host=localhost port=5432 user=postgres dbname=MemoCards password=Goni sslmode=disable"

func ConnectionDb() {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()
}

func InitialMigration() {
	//ConnectionDb()
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	db.AutoMigrate(&Deck{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Commentaire{})

}
