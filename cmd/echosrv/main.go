package main

import (
	"butterfly.orx.me/core"
	"butterfly.orx.me/core/app"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi/operations"
	"go.orx.me/echosrv/internal/db"
	"go.orx.me/echosrv/internal/handler"
	"go.orx.me/echosrv/internal/object"
	"go.orx.me/echosrv/internal/profiler"
)

func main() {
	go profiler.Init()
	db.Init()
	object.Init()
	app := core.New(&app.Config{
		Service: "echo",
		Router:  handler.Router,
	})
	app.Run()
}
