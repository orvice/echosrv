package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/otel"
	"go.orx.me/echosrv/ent"
	"go.orx.me/echosrv/ent/accesslog"
	"go.orx.me/echosrv/internal/db"
	"go.orx.me/echosrv/internal/object"
)

func Router(r *gin.Engine) {
	r.Use(loggingMiddleware)
	r.GET("/ping", Ping)
	r.GET("/healthz", Ping)
	r.GET("/asc/:text", ASC)
	oauthRouter(r)
}

type accessLog struct {
	Data   *ent.AccessLog
	Header http.Header
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

	b, err := json.Marshal(accessLog{
		Data:   log,
		Header: c.Request.Header,
	})
	if err != nil {
		slog.Error("failed to marshal access log", "error", err)
		return
	}

	buf := bytes.NewReader(b)
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	ojbectName := fmt.Sprintf("accesslogs/%s/%s.json", time.Now().Format("2006/01/02"), uuid)

	info, err := object.Client.PutObject(c.Request.Context(), bucketName, ojbectName, buf, int64(len(b)), minio.PutObjectOptions{})
	if err != nil {
		slog.Error("failed to put object", "error", err)
		return
	}

	slog.Info("access log",
		slog.Time("start", start), slog.Time("end", time.Now()),
		slog.String("trace", trace),
		slog.Int("id", log.ID),
		slog.String("client_ip", c.ClientIP()),
		slog.String("uuid", uuid),
		slog.String("object.key", info.Key),
		slog.String("object.etag", info.ETag),
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

	t := time.Now().Unix()
	if t%9 == 0 {
		go func() {
			// clean old data
			rows, err := cli.AccessLog.Delete().
				Where(accesslog.CreatedUnixLT(int(t - 86400))).
				Exec(context.Background())
			if err != nil {
				slog.Error("failed to delete old access log", "error", err)
			} else {
				slog.Info("delete old access log", slog.Int("rows", rows))
			}
		}()
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"headers": c.Request.Header,
	})
}

func ASC(c *gin.Context) {
	myFigure := figure.NewFigure(c.Param("text"), "", true)
	c.String(200, myFigure.String())
}

func users() {
	cli := db.EntClient()
	users, err := cli.User.Query().All(context.Background())
	if err != nil {
		slog.Error("failed to query users", "error", err)
	}
	slog.Info("users", slog.Int("users_count", len(users)))
}
