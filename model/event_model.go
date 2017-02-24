package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Event is the structure defining the event properties
type Event struct {
	gorm.Model           //`json:"-"`
	Name       string    `json:"name"`
	Place      string    `json:"place"`
	Begin      time.Time `json:"begin"`
	End        time.Time `json:"end"`
	Stat       EventStat `json:"stat"`
}

// NewEvent create a event which is the main data structure for stats
func NewEvent() *Event {
	return &Event{}
}
