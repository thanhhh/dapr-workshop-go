package main

import (
	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/server"
	"dapr-workshop-go/pkg/utils"
	"log"
	"os"

	"github.com/labstack/echo/v4"

	fcServer "dapr-workshop-go/fine-collection-service/internal/server"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
		cfg.Server.SSL)

	echo := echo.New()

	serverHandlers := fcServer.NewServerHandler(echo, cfg, appLogger)

	s := server.NewServer(echo, cfg, appLogger, serverHandlers)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
