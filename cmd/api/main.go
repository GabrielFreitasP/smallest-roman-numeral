package main

import (
	"log"

	"github.com/GabrielFreitasP/smallest-roman-numeral/config"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/server"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
)

const configPath = "./config/config-local"

// @title Smallest Roman Numeral API
// @version 1.0
// @description Smallest Roman Numeral API
// @contact.name Gabriel de Freitas Pinheiro
// @contact.url https://github.com/GabrielFreitasP
// @contact.email gabrieldefreitaspinheiro@gmail.com
func main() {
	log.Println("Starting API server")

	// Init configuration
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	// Init logger
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	// Start server
	s := server.NewServer(cfg, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
