package server

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	db "github.com/chaitanyabsprip/journal/pkg/database"
)

type Entry struct {
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	Detail    string    `json:"detail"`
}

func InsertEntry(ctx context.Context, entry *Entry) {
	db.Client.Database("journal").Collection("entries").InsertOne(ctx, entry)
}

func GetEnteries(ctx context.Context) ([]Entry, error) {
	coll := db.Client.Database("journal").Collection("entries")
	curr, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	entries := make([]Entry, 0)
	err = curr.All(ctx, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
