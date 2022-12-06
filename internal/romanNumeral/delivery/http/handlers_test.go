package http

import (
	"errors"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/logger"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral/mock"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/converter"
	"github.com/GabrielFreitasP/smallest-roman-numeral/pkg/httpErrors"
)

func TestRomanNumeralHandlers_Search(t *testing.T) {
	cases := []struct {
		name         string
		param        *models.RomanNumeralSearch
		setMocks     func(*mock.MockUseCase, ...interface{})
		expectedCode int
		expectedRes  interface{}
	}{
		{
			name:         "it should respond bad request",
			param:        nil,
			setMocks:     func(_ *mock.MockUseCase, _ ...interface{}) {},
			expectedCode: http.StatusBadRequest,
			expectedRes: &httpErrors.RestError{
				ErrStatus: http.StatusBadRequest,
				ErrError:  httpErrors.BadRequest.Error(),
				ErrCauses: errors.New("EOF"),
			},
		},
		{
			name:  "it should not found",
			param: &models.RomanNumeralSearch{Text: "I"},
			setMocks: func(mockRomanNumeralUC *mock.MockUseCase, models ...interface{}) {
				mockRomanNumeralUC.
					EXPECT().
					Search(gomock.Any(), models[0]).
					Return(nil, romanNumeral.PrimeRomanNumeralNotFound)
			},
			expectedCode: http.StatusNotFound,
			expectedRes: &httpErrors.RestError{
				ErrStatus: http.StatusNotFound,
				ErrError:  romanNumeral.PrimeRomanNumeralNotFound.Error(),
				ErrCauses: romanNumeral.PrimeRomanNumeralNotFound,
			},
		},
		{
			name:  "it should internal server error",
			param: &models.RomanNumeralSearch{Text: "I"},
			setMocks: func(mockRomanNumeralUC *mock.MockUseCase, models ...interface{}) {
				mockRomanNumeralUC.
					EXPECT().
					Search(gomock.Any(), models[0]).
					Return(nil, httpErrors.InternalServerError)
			},
			expectedCode: http.StatusInternalServerError,
			expectedRes: &httpErrors.RestError{
				ErrStatus: http.StatusInternalServerError,
				ErrError:  httpErrors.InternalServerError.Error(),
				ErrCauses: httpErrors.InternalServerError,
			},
		},
		{
			name:  "it should respond OK",
			param: &models.RomanNumeralSearch{Text: "II"},
			setMocks: func(mockRomanNumeralUC *mock.MockUseCase, models ...interface{}) {
				mockRomanNumeralUC.
					EXPECT().
					Search(gomock.Any(), models[0]).
					Return(models[1], nil)
			},
			expectedCode: http.StatusOK,
			expectedRes: &models.RomanNumeral{
				Number: "II",
				Value:  2,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRomanNumeralUC := mock.NewMockUseCase(ctrl)
			tc.setMocks(mockRomanNumeralUC, tc.param, tc.expectedRes)

			h := NewRomanNumeralHandlers(mockRomanNumeralUC, logger.NewFmtLogger())
			handlerFunc := h.Search()

			// If param is nil, the request body going to be nil too
			var body io.Reader = nil
			if tc.param != nil {
				buf, _ := converter.AnyToBytesBuffer(tc.param)
				body = strings.NewReader(buf.String())
			}

			req := httptest.NewRequest(http.MethodPost, "/search", body)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			res := httptest.NewRecorder()
			res.Header().Set(echo.HeaderXRequestID, uuid.New().String())

			ctx := echo.New().NewContext(req, res)

			// Act
			err := handlerFunc(ctx)

			// Assert
			assert.Nil(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, tc.expectedCode, res.Code)

			expectedBuf, _ := converter.AnyToBytesBuffer(tc.expectedRes)
			assert.NotNil(t, expectedBuf)
			assert.Equal(t, expectedBuf, res.Body)
		})
	}
}
