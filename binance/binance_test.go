package binance_test

import (
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/binance"
	"github.com/jeremiasbots/currencyRatesLib/utils"
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
	_, err := binance.GetDollarValueAsFloat64()
	if err != nil {
		t.Error(err)
	}
}

func TestDollarValueAsFloat32(t *testing.T) {
	_, err := binance.GetDollarValueAsFloat32()
	if err != nil {
		t.Error(err)
	}
}

func TestData(t *testing.T) {
	_, err := binance.GetDollarData()
	if err != nil {
		t.Error(err)
	}
}
