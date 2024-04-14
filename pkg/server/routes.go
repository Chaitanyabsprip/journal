package server

import (
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", handleIndex())
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("POST /api/v1/entry", handlePostEntry())
}
