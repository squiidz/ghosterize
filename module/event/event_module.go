package event

import (
	"net/http"

	"local/ghosterize/handler/event"
	"local/ghosterize/model"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

type Event struct {
	db *gorm.DB
}

func New(db *gorm.DB) Event {
	return Event{db}
}

// Module iniate the event module
func (e Event) Module() http.Handler {
	mux := bone.New()
	mux.GetFunc("/show", event.ShowAll)
	mux.GetFunc("/show/:id", event.Show)
	mux.GetFunc("/show/:id/invitations", event.ShowInvitations)
	mux.PostFunc("/create", event.Create)
	mux.PostFunc("/delete/:id", event.Delete)
	e.initSchema()
	return mux
}

func (e Event) initSchema() {
	if !e.db.HasTable(&model.Event{}) {
		e.db.Table("events")
	}
	e.db.AutoMigrate(&model.Event{})
}
