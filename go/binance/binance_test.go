package binance_test

import (
	"fmt"
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/go/binance"
	"github.com/jeremiasbots/currencyRatesLib/go/utils"
)

func TestStringAndBasicValue(t *testing.T) {
	dollarValue, err := binance.GetDollarValue()
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(dollarValue) {
		t.Errorf("The value does not have eight decimal places: %s", dollarValue)
	}
}

func TestDollarValueAsFloat64(t *testing.T) {
	dollarValue, err := binance.GetDollarValueAsFloat64()
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(fmt.Sprintf("%.8f", dollarValue)) {
		t.Errorf("The value does not have eight decimal places: %.8f", dollarValue)
	}
}

func TestData(t *testing.T) {
	_, err := binance.GetDollarData()
	if err != nil {
		t.Error(err)
	}
}
