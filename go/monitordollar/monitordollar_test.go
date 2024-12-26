package monitordollar_test

import (
	"fmt"
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/go/monitordollar"
	"github.com/jeremiasbots/currencyRatesLib/go/utils"
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
	value, err := monitordollar.GetDollarValueAsFloat64()
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(fmt.Sprintf("%.8f", value)) {
		t.Errorf("The value does not have eight decimal places: %.8f", value)
	}
}

func TestDollarData(t *testing.T) {
	_, err := monitordollar.GetDollarData()
	if err != nil {
		t.Error(err)
	}
}
