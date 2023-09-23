package db

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func Ping(ctx context.Context) {
	db, err := otelsql.Open("mysql", os.Getenv("DSN"),
		otelsql.WithAttributes(semconv.DBSystemSqlite),
		otelsql.WithDBName("mydb"),
	)
	if err != nil {
		slog.Error("failed to connect", "error", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		slog.Error("failed to ping ", "error", err)
	}

	slog.Info("Successfully connected to DB!")
}
