package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))
	slog.LogAttrs(context.Background(), slog.LevelInfo, "",
		slog.Group("abc", slog.String("a", "b"), slog.String("c", "d")))
}
