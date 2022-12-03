package middleware

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/sanitize"
)

// Sanitize and read request body to ctx for next use in easy json
func (mw *Manager) Sanitize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(ctx.Request().Body)

		sanBody, err := sanitize.SanitizeJSON(body)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		ctx.Set("body", sanBody)
		return next(ctx)
	}
}
