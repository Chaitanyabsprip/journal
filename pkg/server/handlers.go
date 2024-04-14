package server

import (
	"encoding/json"
	"net/http"
	"text/template"
)

var indexTmpl *template.Template = template.Must(template.ParseFiles("views/index.html"))

func handlePostEntry() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var entry Entry
			err := json.NewDecoder(r.Body).Decode(&entry)
			if err != nil {
				http.Error(w, "Failed to parse request body", http.StatusBadRequest)
				return
			}
			InsertEntry(r.Context(), &entry)
			entries, err := GetEnteries(r.Context())
			if err != nil {
				http.Error(w, "Failed to retrieve log entries", http.StatusInternalServerError)
				return
			}
			indexTmpl.ExecuteTemplate(w, "log", entries)
		},
	)
}

func handleIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		entries, err := GetEnteries(r.Context())
		if err != nil {
			http.Error(w, "Failed to retrieve log entries", http.StatusInternalServerError)
			return
		}
		indexTmpl.ExecuteTemplate(w, "index", entries)
	})
}
