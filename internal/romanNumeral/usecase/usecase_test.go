package usecase

import (
	"context"
	"testing"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
)

func TestUseCase_Search(t *testing.T) {
	t.Parallel()

	uc := NewRomanNumeralUseCase()

	romanNumeralSearch := &models.RomanNumeralSearch{
		Text: "AA",
	}

	_, _ = uc.Search(context.Background(), romanNumeralSearch)
	//res, err := uc.Search(context.Background(), romanNumeralSearch)

	//assert.NoError(t, err)
	//assert.NotNil(t, res)
	//assert.Equal(t, "XI", res.Value)
	//assert.Equal(t, 11, res.Value)
}
