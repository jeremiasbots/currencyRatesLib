package bcv_test

import (
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/go/bcv"
	"github.com/jeremiasbots/currencyRatesLib/go/utils"
)

func TestEuroValue(t *testing.T) {
	value, err := bcv.GetCurrencyValue(bcv.Euro)
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}

func TestYuanValue(t *testing.T) {
	value, err := bcv.GetCurrencyValue(bcv.Yuan)
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}

func TestTurkishLiraValue(t *testing.T) {
	value, err := bcv.GetCurrencyValue(bcv.TurkishLira)
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}

func TestRubleValue(t *testing.T) {
	value, err := bcv.GetCurrencyValue(bcv.Ruble)
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}

func TestAmericanDollarValue(t *testing.T) {
	value, err := bcv.GetCurrencyValue(bcv.AmericanDollar)
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}
