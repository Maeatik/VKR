package main

import (
	"diploma/internal/server"
	log "diploma/pkg/logger"
	"fmt"

	"diploma/config"
	"os"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	logger := log.New("debug","N/A")
	logger.Info(">> Starting service")

	cfg, port, err := initConfig()
	if err != nil{
		logger.Fatal(fmt.Sprintf("error initializating configs: %s", err.Error()))
	}
	
	if err := server.Run(cfg, port, logger); err != nil{
		logger.Fatal(fmt.Sprintf("FATAL: %s",err.Error()))
	}

	logger.Info(">> Stopping application...")
}

func initConfig() (config.Config, string, error)  {
	var cfg config.Config
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	
	if err := gotenv.Load(); err != nil{
		return config.Config{}, "", fmt.Errorf("error initializating env file: %s", err.Error())
	}
	err := viper.ReadInConfig();
	
	if err != nil{
		cfg.Host = viper.GetString("db.host")
		cfg.Port = viper.GetString("db.port")
		cfg.Username = viper.GetString("db.user")
		cfg.Password = os.Getenv("DB_PASSWORD")
		cfg.DBName = viper.GetString("db.dbname")
	}

	return cfg, viper.GetString("port"), err
}
