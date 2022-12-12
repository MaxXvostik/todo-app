package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/maXvostik/todo-app"
	"github.com/maXvostik/todo-app/pkg/handler"
	"github.com/maXvostik/todo-app/pkg/repositori"
	"github.com/maXvostik/todo-app/pkg/service"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initilizing configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatalf("error loading env variabls: %s", err.Error())
	}

	db, err := repositori.NewPostgresDB(repositori.Config{
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialized db: %s", err.Error())
	}

	repos := repositori.NewRepositori(db)
	services := service.NewService(repos)
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
