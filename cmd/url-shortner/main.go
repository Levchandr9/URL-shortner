package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"url-shortner/config/pkg/config"
	"url-shortner/config/pkg/lib/logger/sl"

	_ "github.com/lib/pq"
)

const (
	envLocal   = "local"
	envDev     = "dev"
	envProd    = "prod"
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "admin"
	dbPassword = "123"
	dbName     = "URL"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortner", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	log.Info("configuration loaded successfully")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error("failed to ping database", sl.Err(err))
		os.Exit(1)
	}

	log.Info("successfully connected to database")

	// TODO: init outer: chi, "chi render"
	log.Info("initialization complete, ready to start server")

	// TODO: run server
	log.Info("server is running")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
