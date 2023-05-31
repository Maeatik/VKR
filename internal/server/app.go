package server

import (
	"diploma/config"
	"diploma/internal/handler"
	"diploma/internal/repository"
	"diploma/internal/service"
	"diploma/pkg/logger"
	"diploma/pkg/postgres"
	"fmt"
)

func Run(cfg config.Config, port string, log *logger.Logger) error{
	serviceDb, err := postgres.New(cfg, log)
	if err != nil {
		log.Fatal("Unable to establish a connection with service database.")
		return err
	}
	defer serviceDb.Close()

	server := new(Server)

	rep := repository.NewRepository(serviceDb)
	services := service.NewServices(rep)
	handlers := handler.NewHandler(*services, log)

	if err := server.Run(port, handlers.InitRoutes()); err != nil{
		fmt.Println(err)
	}

	return nil
}