package handler

import (
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Home"))
}
