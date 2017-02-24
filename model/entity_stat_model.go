package model

import "github.com/jinzhu/gorm"

// EntityStat definied the properties of the entity statistics
type EntityStat struct {
	gorm.Model
	Invited      []*Event `json:"invited"`
	Participated []*Event `json:"participated"`
	Global       int      `json:"global"`
}

// NewEntityStat provide the default stat values and return a new stat struct
func NewEntityStat() *EntityStat {
	return &EntityStat{
		Invited:      []*Event{},
		Participated: []*Event{},
		Global:       0,
	}
}
