package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral"
)

func TestRomanNumeralUseCase_Search(t *testing.T) {
	cases := []struct {
		name        string
		param       *models.RomanNumeralSearch
		expectedRes *models.RomanNumeral
		expectedErr error
	}{
		{
			name:        "it should return not found error when text is empty",
			param:       &models.RomanNumeralSearch{Text: ""},
			expectedRes: nil,
			expectedErr: romanNumeral.PrimeRomanNumeralNotFound,
		},
		{
			name:        "it should return not found error when not found roman numeral",
			param:       &models.RomanNumeralSearch{Text: "AABB"},
			expectedRes: nil,
			expectedErr: romanNumeral.PrimeRomanNumeralNotFound,
		},
		{
			name:        "it should return not found error when not found prime roman numeral",
			param:       &models.RomanNumeralSearch{Text: "AAXBB"},
			expectedRes: nil,
			expectedErr: romanNumeral.PrimeRomanNumeralNotFound,
		},
		{
			name:        "it should return the roman numeral II",
			param:       &models.RomanNumeralSearch{Text: "AAXBBIVXIBII"},
			expectedRes: &models.RomanNumeral{Number: "II", Value: 2},
			expectedErr: nil,
		},
		{
			name:        "it should return the roman numeral III",
			param:       &models.RomanNumeralSearch{Text: "AAXBBIIIVXIBV"},
			expectedRes: &models.RomanNumeral{Number: "III", Value: 3},
			expectedErr: nil,
		},
		{
			name:        "it should return the roman numeral LXI",
			param:       &models.RomanNumeralSearch{Text: "MMMJDDURIROGOFJHAIXIXIJFJJLLSLXIJJUCYIDJFEJ"},
			expectedRes: &models.RomanNumeral{Number: "LXI", Value: 61},
			expectedErr: nil,
		},
		{
			name:        "it should return the roman numeral MLXIII",
			param:       &models.RomanNumeralSearch{Text: "AABBIAAIJFIAHFIEHFHISJOMLXIIIAOFKOAFHNCNGMNXBHDFAPIDFJI"},
			expectedRes: &models.RomanNumeral{Number: "MLXIII", Value: 1063},
			expectedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()

			uc := NewRomanNumeralUseCase()

			// Act
			res, err := uc.Search(ctx, tc.param)

			// Assert
			if tc.expectedRes == nil {
				assert.Nil(t, res)
			} else {
				assert.NotNil(t, res)
				assert.Equal(t, tc.expectedRes, res)
			}

			if tc.expectedErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedErr, err)
			}
		})
	}
}
