package main

import (
	"context"
	"log/slog"

	"butterfly.orx.me/core"
	"butterfly.orx.me/core/app"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi/operations"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.orx.me/echosrv/internal/db"
	"go.orx.me/echosrv/internal/handler"
	"go.orx.me/echosrv/internal/object"
	"go.orx.me/echosrv/internal/profiler"
	"go.uber.org/fx"
	"google.golang.org/protobuf/encoding/protojson"
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
			slog.Info("app run")
			go app.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("stop app")
			return nil
		},
	})
	return app
}

func main() {

	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
		}),
	)

	fx.New(
		fx.Provide(NewApp),
		fx.Invoke(func(*app.App) {}),
	).Run()
}
