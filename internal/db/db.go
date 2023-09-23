package db

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	db *sql.DB
)

func Init() {
	var err error
	db, err = otelsql.Open("mysql", os.Getenv("DSN"),
		otelsql.WithAttributes(semconv.DBSystemSqlite),
		otelsql.WithDBName("mydb"),
	)
	if err != nil {
		slog.Error("failed to connect", "error", err)
	}
}

func Ping(ctx context.Context) {
	if err := db.PingContext(ctx); err != nil {
		slog.Error("failed to ping ", "error", err)
		return
	}
	slog.Info("Successfully ping  to DB!")
}
