package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	apiMiddleware "github.com/GabrielFreitasP/smallest-roman-numeral/internal/middleware"
	romanNumeralHttp "github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral/delivery/http"
	romanNumeralUseCase "github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral/usecase"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/metric"
)

// Map handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	// Metrics
	metrics, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("Create metrics error: %s", err)
	} else {
		s.logger.Infof("Metrics available URL: %s, ServiceName: %s", s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	}

	// UseCases
	numUC := romanNumeralUseCase.NewSessionUseCase()

	// Handlers
	numHandlers := romanNumeralHttp.NewRomanNumeralHandlers(numUC, s.logger)

	// Middlewares
	mw := apiMiddleware.NewMiddlewareManager(s.logger)
	e.Use(mw.RequestLoggerMiddleware)
	e.Use(mw.MetricsMiddleware(metrics))
	e.Use(middleware.RequestID())

	// Routes
	romanNumeralHttp.MapRomanNumeralRoutes(e, numHandlers)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
