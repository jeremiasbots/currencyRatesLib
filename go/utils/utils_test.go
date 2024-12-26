package utils_test

import (
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/go/utils"
)

func TestEightDecimals(t *testing.T) {
	eightDecimalsNumberAsString := "56.00000000"
	if !utils.HasEightDecimals(eightDecimalsNumberAsString) {
		t.Errorf("The function is wrong, the value does have eight decimal places: %s", eightDecimalsNumberAsString)
	}
}
