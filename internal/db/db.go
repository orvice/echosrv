package db

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Ping() {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		slog.Error("failed to connect", "error", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("failed to ping ", "error", err)
	}

	slog.Info("Successfully connected to DB!")
}
