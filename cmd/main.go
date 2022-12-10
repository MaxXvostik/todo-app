package main

import (
	"log"

	"github.com/maXvostik/todo-app"
	"github.com/maXvostik/todo-app/pkg/handler"
	"github.com/maXvostik/todo-app/pkg/repositori"
	"github.com/maXvostik/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initilizing configs: %s", err.Error())
	}

	repos := repositori.NewRepositori()
	services := service.NewService(*repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
