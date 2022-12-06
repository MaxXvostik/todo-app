package main

import (
	"log"

	"github.com/maXvostik/todo-app"
	"github.com/maXvostik/todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s ", err.Error())
	}
}
