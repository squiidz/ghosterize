package entity

import (
	"net/http"

	"local/ghosterize/handler/entity"
	"local/ghosterize/model"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

type Entity struct {
	db *gorm.DB
}

func New(db *gorm.DB) Entity {
	return Entity{db}
}

// Module iniate the entitiy module
func (e Entity) Module() http.Handler {
	mux := bone.New()
	mux.GetFunc("/show", entity.ShowAll)
	mux.GetFunc("/show/:id", entity.Show)
	mux.GetFunc("/show/:id/invited", entity.ShowInvited)
	mux.GetFunc("/show/:id/participated", entity.ShowParticipated)
	mux.PostFunc("/create", entity.Create)
	mux.PostFunc("/delete/:id", entity.Delete)
	e.initSchema()
	return mux
}

func (e Entity) initSchema() {
	if !e.db.HasTable(&model.Entity{}) {
		e.db.Table("entities")
		e.db.Table("entity_stats")
	}
	e.db.AutoMigrate(&model.Entity{})
	e.db.AutoMigrate(&model.EntityStat{})
}
