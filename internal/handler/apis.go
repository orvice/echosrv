package handler

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

// Package-level tracer.
// This should be configured in your code setup instead of here.
var tracer = otel.Tracer("handler")

func Ping(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "sleep")
	defer span.End()

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
