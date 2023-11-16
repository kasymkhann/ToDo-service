package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	todo "to-doProjectGo"
	"to-doProjectGo/pkg/handler"
	"to-doProjectGo/pkg/repository"
	"to-doProjectGo/pkg/service"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("Error init configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatalf("err env %s", err.Error())
	}

	db, err := repository.PostgresqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("PASSWORD_db"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error connection db %s", err.Error())
	}

	repo := repository.Repositories(db) // <- db
	service := service.Servic(repo)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)

	go func() {

		if err := srv.Start(viper.GetString("port"), handler.ThisRouter()); err != nil {
			log.Fatalf("Server error: %s", err.Error())

		}
	}()
	log.Println("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("App shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())

	}
	if err := db.Close(); err != nil {
		log.Fatalf("error occured on Database connection close: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
