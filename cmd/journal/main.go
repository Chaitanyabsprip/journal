package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	db "github.com/chaitanyabsprip/journal/pkg/database"
	"github.com/chaitanyabsprip/journal/pkg/server"
)

var port string = "8080"

func run(ctx context.Context, getenv func(string) string) error {
	_, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	_, err := db.Connect(ctx, getenv("MONGODB_URI"))
	if err != nil {
		return err
	}
	defer db.Disconnect(ctx)
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
	srv := server.NewServer(logger)
	fmt.Printf("listening on %v\n", port)
	return http.ListenAndServe(":8080", srv)
}

func main() {
	if err := godotenv.Overload(); err != nil {
		log.Fatal("No .env file found")
	}
	ctx := context.Background()
	if err := run(ctx, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
