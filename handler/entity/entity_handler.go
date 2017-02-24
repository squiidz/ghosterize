package entity

import (
	"encoding/json"
	"net/http"

	db "local/ghosterize/db/sqlite"
	"local/ghosterize/model"

	"github.com/go-zoo/bone"
)

func Show(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	var entity model.Entity
	db.Store.First(&entity, id)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&entity)
}

func ShowAll(rw http.ResponseWriter, req *http.Request) {
	var entities []model.Entity
	db.Store.Find(&entities)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&entities)
}

func ShowInvited(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	var entity model.Entity
	var entityStat model.EntityStat

	db.Store.Find(&entity, id)
	db.Store.Debug().Model(&entityStat).Related(&entity)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&entity.Stat.Invited)
}

func ShowParticipated(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	var entity model.Entity
	db.Store.Find(&entity, id)
	db.Store.Debug().Model(&entity).Related(&entity.Stat.Participated, "stat")
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&entity.Stat.Participated)
}

func Create(rw http.ResponseWriter, req *http.Request) {
	var entity model.Entity
	json.NewDecoder(req.Body).Decode(&entity)
	db.Store.Create(&entity)
}

func Delete(rw http.ResponseWriter, req *http.Request) {

}
