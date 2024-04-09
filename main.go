package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/chaitanyabsprip/journal/database"
)

func main() {
	db.Connect()
	defer db.Disconnect()
	http.HandleFunc("POST /entry", handlePostEntry)
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePostEntry(w http.ResponseWriter, r *http.Request) {
	var entry Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	InsertEntry(&entry)
	w.WriteHeader(http.StatusOK)
}
