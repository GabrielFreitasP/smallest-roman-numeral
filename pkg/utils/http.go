package utils

import (
	"context"
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"

	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/httpErrors"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/sanitize"
)

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get context  with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

// Error response with logging error for echo context
func ErrResponseWithLog(ctx echo.Context, logger logger.Logger, err error) error {
	LogResponseError(ctx, logger, err)
	return ctx.JSON(httpErrors.ErrorResponse(err))
}

// Error response with logging error for echo context
func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}

// Read sanitize and validate request
func SanitizeRequest(ctx echo.Context, request interface{}) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(ctx.Request().Body)

	sanBody, err := sanitize.SanitizeJSON(body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(sanBody, request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}
