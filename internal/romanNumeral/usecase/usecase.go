package usecase

import (
	"context"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral"
	"github.com/opentracing/opentracing-go"
	"math"
	"regexp"
)

const RomanNumeralPattern = `M{0,4}(C[MD]|D?C{0,3})(X[CL]|L?X{0,3})(I[XV]|V?I{0,3})`

// Roman numerals
var RomanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// Roman numeral use case struct
type RomanNumeralUseCase struct {
}

// Roman numeral use case constructor
func NewRomanNumeralUseCase() *RomanNumeralUseCase {
	return &RomanNumeralUseCase{}
}

// Search the smaller roman numeral in text
func (uc *RomanNumeralUseCase) Search(ctx context.Context, romanNumSearch *models.RomanNumeralSearch) (*models.RomanNumeral, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "numUC.Search")
	defer span.Finish()

	romanNums := uc.getRomanNumbers(romanNumSearch.Text)

	var romanNumNumber string
	var romanNumValue int
	for _, num := range romanNums {
		if len(num) == 0 {
			continue
		}
		value := uc.romanToInt(num)
		isPrime := uc.isPrimeNumber(value)
		if isPrime && (romanNumValue == 0 || value < romanNumValue) {
			romanNumNumber = num
			romanNumValue = value
		}
	}

	if romanNumValue == 0 {
		return nil, romanNumeral.PrimeRomanNumeralNotFound
	}

	romanNum := &models.RomanNumeral{
		Number: romanNumNumber,
		Value:  romanNumValue,
	}

	return romanNum, nil
}

// Get roman numeral from a text
func (uc *RomanNumeralUseCase) getRomanNumbers(s string) []string {
	r := regexp.MustCompile(RomanNumeralPattern)
	return r.FindAllString(s, -1)
}

// Parse roman numeral to int value
func (uc *RomanNumeralUseCase) romanToInt(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}

// Check if the number is prime
func (uc *RomanNumeralUseCase) isPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}

	sq := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
