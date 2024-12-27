package utils_test

import (
	"errors"
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/go/utils"
)

func TestEightDecimals(t *testing.T) {
	eightDecimalsNumberAsString := "56.00000000"
	if !utils.HasEightDecimals(eightDecimalsNumberAsString) {
		t.Errorf("The function is wrong, the value does have eight decimal places: %s", eightDecimalsNumberAsString)
	}
}

func TestErrAsNil(t *testing.T) {
	firstErr := utils.ErrAsNil(nil)
	if firstErr != "null" {
		t.Error("Error isn´t \"null\"")
	}
	secondErr := utils.ErrAsNil(errors.New("example"))
	if secondErr != "example" {
		t.Error("Error isn´t \"example\"")
	}
}
