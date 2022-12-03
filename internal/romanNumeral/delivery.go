package romanNumeral

import "github.com/labstack/echo/v4"

// Search HTTP handlers interface
type Handlers interface {
	Search() echo.HandlerFunc
}
