package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestMain(t *testing.T) {
	TestDataBaseInit()

}
func TestNewUser(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
	}
	req := httptest.NewRequest("PUT", "/user/TestUser1/test1@gmail.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	t.Errorf("Database have 1 records of user ")

}
func TestAllUser(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
	}
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

}

/*func TestFindUserById(t *testing.T) {
	user := &User{Name: "TestUser2", Prenom: "TestPrenom2", Email: "test2@gmail.com"}

	CreateUser(&User)
	carFound := GetUser()

	if car.Id != carFound.Id {
		t.Error("Couldn't find car by id")
	}
}*/
