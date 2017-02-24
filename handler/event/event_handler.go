package event

import (
	"encoding/json"
	"net/http"

	db "local/ghosterize/db/sqlite"
	"local/ghosterize/model"

	"github.com/go-zoo/bone"
)

func Show(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	var event model.Event
	db.Store.First(&event, id)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&event)
}

func ShowAll(rw http.ResponseWriter, req *http.Request) {
	var events []model.Event
	err := db.Store.Find(&events)
	if err != nil {
		rw.Write([]byte("Invalid data"))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&events)
}

func ShowInvitations(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	var event model.Event
	db.Store.Find(&event, id)
	db.Store.Debug().Model(&event).Related(&event.Stat.Invited, "invitations")
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&event.Stat.Invited)
}

func Create(rw http.ResponseWriter, req *http.Request) {
	var evt model.Event
	err := json.NewDecoder(req.Body).Decode(&evt)
	if err != nil {
		rw.Write([]byte("Invalid request data"))
	}
	db.Store.Create(&evt)
}

func Delete(rw http.ResponseWriter, req *http.Request) {

}
