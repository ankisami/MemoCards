package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Commentaire struct {
	gorm.Model
	Note        string
	Description string
	Date        string
	DeckID      uint
	UserID      uint
}
