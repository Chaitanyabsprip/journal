package main

import (
	"context"

	db "github.com/chaitanyabsprip/journal/database"
)

type Entry struct {
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

func InsertEntry(entry *Entry) {
	client := db.Client
	coll := client.Database("journal").Collection("entries")
	coll.InsertOne(context.TODO(), entry)
}
