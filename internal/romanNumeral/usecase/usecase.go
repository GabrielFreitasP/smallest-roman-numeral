package usecase

import (
	"context"
	"fmt"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/models"
	"github.com/GabrielFreitasP/smallest-roman-numeral/internal/romanNumeral"
	"github.com/opentracing/opentracing-go"
	"math"
	"regexp"
)

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
type UseCase struct {
}

// Roman numeral use case constructor
func NewRomanNumeralUseCase() *UseCase {
	return &UseCase{}
}

// Search the smaller roman numeral in text
func (uc *UseCase) Search(ctx context.Context, romanNumSearch *models.RomanNumeralSearch) (*models.RomanNumeral, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "numUC.Search")
	defer span.Finish()

	text := romanNumSearch.Text
	fmt.Println(text)

	pattern := `M{0,4}(C[MD]|D?C{0,3})(X[CL]|L?X{0,3})(I[XV]|V?I{0,3})`
	r := regexp.MustCompile(pattern)
	romanNums := r.FindAllString(text, -1)

	var romanNumNumber string
	var romanNumValue int
	for _, num := range romanNums {
		if len(num) > 0 {
			value := romanToInt(num)
			isPrime := isPrimeNumber(value)
			fmt.Printf("%s|%d|%t\n", num, value, isPrime)
			if isPrime && (romanNumValue == 0 || value < romanNumValue) {
				romanNumNumber = num
				romanNumValue = value
			}
		}
	}

	if romanNumValue == 0 {
		return nil, romanNumeral.PrimeRomanNumeralNotFound
	}

	romanNum := &models.RomanNumeral{
		Number: romanNumNumber,
		Value:  romanNumValue,
	}

	fmt.Println(romanNum)

	return romanNum, nil
}

func romanToInt(s string) int {
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

func isPrimeNumber(num int) bool {
	if num < 2 {
		return true
	}

	sq := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
