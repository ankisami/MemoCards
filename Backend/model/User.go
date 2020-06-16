package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}
type error interface {
	Error() string
}

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

	user := &User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)
	db.Create(user)

	/*vars := mux.Vars(r)
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
	}*/

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

func Login(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOne(email, password string) map[string]interface{} {
	user := &User{}

	/*db.Where("email= ? ", email).Find(&user)
	if user.ID == 0 {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}*/
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
