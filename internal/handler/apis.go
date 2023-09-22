package handler

import (
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.orx.me/echosrv/internal/db"
	"golang.org/x/exp/slog"
)

func Router(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.GET("/healthz", Ping)
	r.GET("/asc/:text", ASC)
}

// Package-level tracer.
// This should be configured in your code setup instead of here.
var tracer = otel.Tracer("handler")

func Ping(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "sleep")
	defer span.End()
	slog.Info("ping", slog.Time("start", time.Now()))
	go db.Ping()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ASC(c *gin.Context) {
	myFigure := figure.NewFigure(c.Param("text"), "", true)
	c.String(200, myFigure.String())
}
