package main

import (
	"log"

	"github.com/maXvostik/todo-app"
	"github.com/maXvostik/todo-app/pkg/handler"
	"github.com/maXvostik/todo-app/pkg/repositori"
	"github.com/maXvostik/todo-app/pkg/service"
)

func main() {

	repos := repositori.NewRepositori()
	services := service.NewService(*repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s ", err.Error())
	}
}
