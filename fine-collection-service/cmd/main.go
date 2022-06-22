package main

import (
	"context"
	"log"
	"os"

	dapr "github.com/dapr/go-sdk/client"
	echo "github.com/labstack/echo/v4"

	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/server"
	"dapr-workshop-go/pkg/utils"

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

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	opt := map[string]string{
    "version": "1",
	}
	ctx := context.Background()
	secret, err := client.GetSecret(ctx, "trafficcontrol-secrets", "finecalculator.licensekey", opt)

	if err != nil {
		panic(err)
	}
	
	cfg.LicenseKey.FineCalculatorLicenseKey = secret["finecalculator.licensekey"]

	e := echo.New()

	serverHandlers := fcServer.NewServerHandler(e, cfg, appLogger)

	s := server.NewServer(e, cfg, appLogger, serverHandlers)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
