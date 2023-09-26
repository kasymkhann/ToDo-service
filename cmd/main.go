package main

import (
	"log"
	todo "to-doProjectGo"
	"to-doProjectGo/pkg/handler"
	"to-doProjectGo/pkg/repository"
	"to-doProjectGo/pkg/service"

	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("Error init configs: %s", err.Error())
	}

	repo := repository.Repositories()
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
