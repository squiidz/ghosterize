package model

import (
	"github.com/jinzhu/gorm"
)

// Entity is the structure defining the user activities
type Entity struct {
	gorm.Model             //`json:"-"`
	Email      string      `json:"email"`
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	FirstName  string      `json:"firstname"`
	LastName   string      `json:"lastname"`
	Stat       *EntityStat `json:"stat" gorm:"ForeignKey:StatID"`
}

// NewEntity create a entity which is the main data structure for stats
func NewEntity(email, password, first, last string) *Entity {
	return &Entity{
		Email:     email,
		Password:  password,
		FirstName: first,
		LastName:  last,
		Stat:      &EntityStat{},
	}
}
