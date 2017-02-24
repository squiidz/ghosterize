package model

import "github.com/jinzhu/gorm"

// EventStat definied the properties of the entity statistics
type EventStat struct {
	gorm.Model   `json:"-"`
	Invited      []*Entity `json:"invited"`
	Participated []*Entity `json:"participated"`
	Global       int       `json:"global"`
}

// NewEventStat provide the default stat values and return a new stat struct
func NewEventStat() *EventStat {
	return &EventStat{
		Invited:      []*Entity{},
		Participated: []*Entity{},
		Global:       0,
	}
}
