package main

import (
	"log"
	"os"
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
	if err := srv.Start(viper.GetString("port"), handler.ThisRouter()); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
