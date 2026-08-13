package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"dvdrental-api/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		logger.Info("Couldn't load .env file.", "error", err.Error())
	}

	cfg := config{
		addr: utils.GetString("PORT", "8000"),
		db: dbConfig{
			dsn: utils.GetString("DB_URL", "postgres://postgres:postgres@localhost:15432/dvdrental?sslmode=disable"),
		},
		rdb: redisConfig{
			redisURL:   utils.GetString("REDIS_URL", "redis://:ro0T@localhost:16379/0"),
			defaultTTL: utils.GetDuration("REDIS_CACHE_TTL", 5*time.Minute),
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

	// Redis
	redisURL, err := redis.ParseURL(cfg.rdb.redisURL)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(redisURL)
	defer rdb.Close()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	logger.Info("redis connection established")

	api := application{
		config: cfg,
		db:     conn,
		rdb:    rdb,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("failed to start a server.", "error", err)
		os.Exit(1)
	}
}
