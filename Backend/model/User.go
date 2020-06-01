package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name     string
	Prenom   string
	Email    string
	Password string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}

	vars := mux.Vars(r)
	email := vars["email"]
	var user User
	db.Where("email= ? ", email).Find(&user)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	password := vars["password"]
	var user User
	db.Where("email= ? ", email).Find(&user)
	if user.ID != 0 {
		fmt.Fprint(w, "Email already exist ")
	} else {
		db.Create(&User{Name: name, Email: email, Password: password})

		fmt.Fprint(w, "NEW USER created ")
	}

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ? ", name).Find(&user)
	db.Delete(&user)

	fmt.Fprint(w, "user deleted")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ? ", name).Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Fprint(w, "user updated")
}
func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	db.Delete(&users)
}

func ConnectionUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNCECT TO DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]
	password := vars["password"]

	var user User
	db.Where("email = ? AND password = ?", email, password).Find(&user)
	if user.ID != 0 {
		fmt.Fprint(w, "user connected")
		json.NewEncoder(w).Encode(user)
	} else {
		fmt.Fprint(w, "Email ou mdp incorrect ")
	}
}
