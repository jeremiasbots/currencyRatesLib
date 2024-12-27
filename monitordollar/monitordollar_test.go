package monitordollar_test

import (
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/monitordollar"
	"github.com/jeremiasbots/currencyRatesLib/utils"
)

func TestDollarAsString(t *testing.T) {
	value, err := monitordollar.GetDollarValue()
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}

func TestDollarAsFloat64(t *testing.T) {
	_, err := monitordollar.GetDollarValueAsFloat64()
	if err != nil {
		t.Error(err)
	}
}

func TestDollarAsFloat32(t *testing.T) {
	_, err := monitordollar.GetDollarValueAsFloat32()
	if err != nil {
		t.Error(err)
	}
}

func TestDollarData(t *testing.T) {
	_, err := monitordollar.GetDollarData()
	if err != nil {
		t.Error(err)
	}
}
