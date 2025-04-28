package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq" // init postgres driver
)

type Storage struct {
	db *sql.DB
}

func New(connStr string) (*Storage, error) {
	fmt.Println("start DB")
	const op = "storage.postgres.New"

	slog.Info("opening database", "connStr", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("failed to open database", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		slog.Error("failed to ping database", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	slog.Info("Executing table creation query")
	slog.Info("creating tables")
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS url(
			id SERIAL PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
		`)
	if err != nil {
		slog.Error("failed to create tables", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	slog.Info("database initialized successfully")
	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
