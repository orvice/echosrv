package main

import (
	"context"

	"butterfly.orx.me/core"
	"butterfly.orx.me/core/app"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi/operations"
	"go.orx.me/echosrv/internal/db"
	"go.orx.me/echosrv/internal/handler"
	"go.orx.me/echosrv/internal/object"
	"go.orx.me/echosrv/internal/profiler"
	"go.uber.org/fx"
)

func NewApp(lc fx.Lifecycle) *app.App {
	go profiler.Init()
	db.Init()
	object.Init()
	app := core.New(&app.Config{
		Service: "echo",
		Router:  handler.Router,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return app
}

func main() {
	fx.New(
		fx.Provide(NewApp),
	).Run()
}
