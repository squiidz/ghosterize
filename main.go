package main

import (
	"fmt"
	"net/http"

	db "local/ghosterize/db/sqlite"
	"local/ghosterize/handler"
	"local/ghosterize/module"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	mw "github.com/go-zoo/claw/middleware"
)

func init() {
	err := db.InitSqlite("data.db") //db.InitPostgres("host=localhost user=postgres dbname=postgres sslmode=disable password=root")
	if err != nil {
		panic(err)
	}
}

func main() {
	cl := claw.New(mw.Logger)
	mux := bone.New()
	modul := module.New(db.Store)
	modul.Mount("/api", "event", mux)
	modul.Mount("/api", "entity", mux)

	mux.GetFunc("/home", handler.HomeHandler)
	mux.GetFunc("/login", handler.LoginHandler)

	fmt.Println("Server listening on :7331")
	http.ListenAndServe(":7331", cl.Merge(mux))
}
