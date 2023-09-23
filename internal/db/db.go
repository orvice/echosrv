package db

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.orx.me/echosrv/ent"
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

	cli := EntClient()
	err = cli.Schema.Create(context.Background())
	if err != nil {
		slog.Error("failed to create schema", "error", err)
	}
}

func EntDrv() *entsql.Driver {
	drv := entsql.OpenDB("mysql", db)
	return drv
}

func EntClient() *ent.Client {
	drv := EntDrv()
	cli := ent.NewClient(ent.Driver(drv))
	return cli
}

func Ping(ctx context.Context) {
	if err := db.PingContext(ctx); err != nil {
		slog.Error("failed to ping ", "error", err)
		return
	}
	slog.Info("Successfully ping  to DB!")
}
