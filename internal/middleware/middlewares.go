package middleware

import (
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
)

// Middleware manager
type Manager struct {
	logger logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(logger logger.Logger) *Manager {
	return &Manager{logger: logger}
}
