package http

import (
	"github.com/labstack/echo/v4"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral"
)

// Map roman numeral routes
func MapRomanNumeralRoutes(e *echo.Echo, h romanNumeral.Handlers) {
	e.POST("/search", h.Search())
}
