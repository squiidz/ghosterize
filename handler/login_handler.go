package handler

import "net/http"

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("login"))
}
