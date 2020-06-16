package model

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Token struct {
	jwt.Claims
	UserID         uint
	Name           string
	Email          string
	StandardClaims jwt.StandardClaims
}
