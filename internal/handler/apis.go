package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	botel "butterfly.orx.me/core/observe/otel"
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.orx.me/echosrv/ent"
	"go.orx.me/echosrv/internal/db"
	"go.orx.me/echosrv/internal/object"
)

var (
	counter metric.Int64Counter
	opt     = metric.WithAttributes(
		attribute.Key("A").String("B"),
		attribute.Key("C").String("D"),
	)
)

func Router(r *gin.Engine) {
	var err error
	counter, err = otel.Meter("api").Int64Counter("foo", metric.WithDescription("a simple counter"))
	if err != nil {
		fmt.Println(err)
	}
	r.Use(loggingMiddleware)
	r.GET("/ping", Ping)
	r.GET("/healthz", Ping)
	r.GET("/asc/:text", ASC)
	r.GET("/metric", prometheusHandler())
	oauthRouter(r)
	r.Any(":any", gin.WrapF(MuxHandler().ServeHTTP))
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.InstrumentMetricHandler(
		botel.PrometheusRegistry(),
		promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
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
	counter.Add(c.Request.Context(), 1, opt)
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

	c.JSON(200, gin.H{
		"message": "pong",
		"headers": c.Request.Header,
	})
}

func ASC(c *gin.Context) {
	myFigure := figure.NewFigure(c.Param("text"), "", true)
	c.String(200, myFigure.String())
}
