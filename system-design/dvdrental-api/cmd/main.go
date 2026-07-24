package main

import (
	"context"
	"log/slog"
	"os"

	"dvdrental-api/internal/env"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		logger.Info("Couldn't load .env file.", "error", err.Error())
	}

	cfg := config{
		addr: env.GetString("PORT", "8000"),
		db: dbConfig{
			dsn: env.GetString("DB_URL", "postgres://postgres:postgres@localhost:15432/dvdrental?sslmode=disable"),
		},
	}

	// Database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("database connection established")

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("failed to start a server.", "error", err)
		os.Exit(1)
	}
}
