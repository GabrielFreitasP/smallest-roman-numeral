package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/utils"
)

// Roman numeral handlers struct
type romanNumeralHandlers struct {
	uc     romanNumeral.UseCase
	logger logger.Logger
}

// Roman numeral handlers constructor
func NewRomanNumeralHandlers(uc romanNumeral.UseCase, logger logger.Logger) *romanNumeralHandlers {
	return &romanNumeralHandlers{uc: uc, logger: logger}
}

// Search
// @Summary Search roman numeral
// @Description Search roman numeral in text
// @Tags RomanNumeral
// @Accept json
// @Produce json
// @Success 200 {object} models.RomanNumeral
// @Failure 404 {object} httpErrors.RestErr
// @Router /search [post]
func (h *romanNumeralHandlers) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "numHandlers.Search")
		defer span.Finish()

		romanNumeralSearch := &models.RomanNumeralSearch{}
		err := utils.SanitizeRequest(c, romanNumeralSearch)
		if err != nil {
			return utils.ErrResponseDefault(c, h.logger, err)
		}

		foundRomanNumeral, err := h.uc.Search(ctx, romanNumeralSearch)
		if err != nil {
			if err == romanNumeral.PrimeRomanNumeralNotFound {
				return utils.ErrResponse(c, h.logger, http.StatusNotFound, err)
			}
			return utils.ErrResponseDefault(c, h.logger, err)
		}

		return c.JSON(http.StatusOK, foundRomanNumeral)
	}
}
