package module

import (
	"errors"
	"net/http"

	"local/ghosterize/module/entity"
	"local/ghosterize/module/event"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

// Router define the interface for a valid Module Router
type Router interface {
	SubRoute(string, bone.Router) *bone.Route
}

// Moduler is defining  the structure of a valid module which could
// be mount to the main application router
type Moduler interface {
	Module() http.Handler
}

var modules = map[string]Moduler{}

// Module is the default construct for the Module system
type Module struct {
	db *gorm.DB
}

// New create a default Module
func New(db *gorm.DB) *Module {
	modules["event"] = event.New(db)
	modules["entity"] = entity.New(db)
	return &Module{db}
}

// Mount mount the provided module by it's name to the passed router
func (m *Module) Mount(prefix, name string, mux Router) error {
	if modules[name] == nil {
		return errors.New("Invalid Module")
	}
	mux.SubRoute(prefix+"/"+name, modules[name].Module())
	return nil
}
