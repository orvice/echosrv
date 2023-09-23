package handler

import (
	"log/slog"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.orx.me/echosrv/internal/db"
)

func Router(r *gin.Engine) {
	r.Use(loggingMiddleware)
	r.GET("/ping", Ping)
	r.GET("/healthz", Ping)
	r.GET("/asc/:text", ASC)
}

func loggingMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	cli := db.EntClient()

	trace := c.GetHeader("Traceparent")
	uuid := uuid.New().String()

	log, err := cli.AccessLog.Create().SetCreatedUnix(int(start.Unix())).SetMethod(c.Request.Method).
		SetPath(c.Request.URL.Path).SetIP(c.ClientIP()).SetUa(uuid).
		SetTrace(trace).
		Save(c.Request.Context())
	if err != nil {
		slog.Error("failed to create access log", "error", err)
		return
	}
	slog.Info("access log",
		slog.Time("start", start), slog.Time("end", time.Now()),
		slog.String("trace", trace),
		slog.Int("id", log.ID),
		slog.String("client_ip", c.ClientIP()),
		slog.String("uuid", uuid),
	)
}

// Package-level tracer.
// This should be configured in your code setup instead of here.
var tracer = otel.Tracer("handler")

func Ping(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "sleep")
	defer span.End()
	slog.Info("ping",
		slog.Time("start", time.Now()),
		slog.String("client_ip", c.ClientIP()),
		slog.String("method", c.Request.Method),
		slog.String("path", c.Request.URL.Path),
		slog.String("user-agent", c.Request.UserAgent()),
	)
	db.Ping(c.Request.Context())

	cli := db.EntClient()
	users, err := cli.User.Query().All(c.Request.Context())
	if err != nil {
		slog.Error("failed to query users", "error", err)
	}
	slog.Info("users", slog.Int("users_count", len(users)))

	c.JSON(200, gin.H{
		"message": "pong",
		"headers": c.Request.Header,
	})
}

func ASC(c *gin.Context) {
	myFigure := figure.NewFigure(c.Param("text"), "", true)
	c.String(200, myFigure.String())
}
