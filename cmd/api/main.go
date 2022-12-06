package main

import (
	"log"

	"github.com/GabrielFreitasP/smallest-roman-numeral/config"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/server"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
)

// @title Smallest Roman Numeral API
// @version 1.0
// @description Smallest Roman Numeral API
// @contact.name Gabriel de Freitas Pinheiro
// @contact.url https://github.com/GabrielFreitasP
// @contact.email gabrieldefreitaspinheiro@gmail.com
func main() {
	cfg := initConfig()
	appLogger := initLogger(cfg)
	initServer(cfg, appLogger)
}

// Init configuration
func initConfig() *config.Config {
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}
	return cfg
}

// Init logger
func initLogger(cfg *config.Config) logger.Logger {
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)
	return appLogger
}

// Init server
func initServer(cfg *config.Config, appLogger logger.Logger) {
	s := server.NewServer(cfg, appLogger)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
