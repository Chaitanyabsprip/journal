package database

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	ctx    context.Context = context.TODO()
)

func Connect() (*mongo.Client, error) {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("invalid environment")
	}
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return Client, nil
}

func Disconnect() error {
	if err := Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
