package main

import (
	"butterfly.orx.me/core"
	"butterfly.orx.me/core/app"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi"
	_ "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi/operations"
	"go.orx.me/echosrv/internal/db"
	"go.orx.me/echosrv/internal/handler"
)

func main() {
	db.Init()
	app := core.New(&app.Config{
		Service: "echo",
		Router:  handler.Router,
	})
	app.Run()
}
