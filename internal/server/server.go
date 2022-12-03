package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/GabrielFreitasP/smallest-roman-numeral/config"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
}

// Server constructor
func NewServer(cfg *config.Config, logger logger.Logger) *Server {
	return &Server{echo: echo.New(), cfg: cfg, logger: logger}
}

// Run server
func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		MaxHeaderBytes: maxHeaderBytes,
	}

	if s.cfg.Server.Mode != "Development" {
		server.ReadTimeout = time.Second * s.cfg.Server.ReadTimeout
		server.WriteTimeout = time.Second * s.cfg.Server.WriteTimeout
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
