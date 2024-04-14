package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(ctx context.Context, dbURI string) (*mongo.Client, error) {
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}
	return Client, nil
}

func Disconnect(ctx context.Context) error {
	if err := Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
