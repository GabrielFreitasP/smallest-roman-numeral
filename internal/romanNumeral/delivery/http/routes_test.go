package http

import (
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapRomanNumeralRoutes(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRomanNumeralHandlers := mock.NewMockHandlers(ctrl)
	mockRomanNumeralHandlers.EXPECT().Search()

	expectedSearchRoute := &echo.Route{
		Method: "POST",
		Path:   "/search",
		Name:   "Search",
	}

	e := echo.New()

	// Act
	MapRomanNumeralRoutes(e, mockRomanNumeralHandlers)

	// Assert
	routes := e.Routes()
	assert.NotNil(t, routes)
	assert.Equal(t, expectedSearchRoute, routes[0])
}
