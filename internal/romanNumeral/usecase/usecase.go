package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
)

// Roman numeral use case struct
type UseCase struct {
}

// Roman numeral use case constructor
func NewSessionUseCase() *UseCase {
	return &UseCase{}
}

// Search the smaller roman numeral in text
func (us *UseCase) Search(ctx context.Context, romanNumeralSearch *models.RomanNumeralSearch) (*models.RomanNumeral, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "numUC.Search")
	defer span.Finish()

	text := romanNumeralSearch.Text

	romanNumeral := &models.RomanNumeral{
		Number: text,
		Value:  len(text),
	}

	return romanNumeral, nil
}
