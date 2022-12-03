package romanNumeral

import (
	"context"

	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
)

// Roman numeral use case interface
type UseCase interface {
	Search(ctx context.Context, romanNumeralSearch *models.RomanNumeralSearch) (*models.RomanNumeral, error)
}
