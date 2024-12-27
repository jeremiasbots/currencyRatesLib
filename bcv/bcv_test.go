package bcv_test

import (
	"testing"

	"github.com/jeremiasbots/currencyRatesLib/bcv"
	"github.com/jeremiasbots/currencyRatesLib/utils"
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

func TestDollarAsFloat64(t *testing.T) {
	_, err := bcv.GetDollarValueAsFloat64()
	if err != nil {
		t.Error(err)
	}
}

func TestDollarValueAsFloat32(t *testing.T) {
	_, err := bcv.GetDollarValueAsFloat32()
	if err != nil {
		t.Error(err)
	}
}

func TestDollarFunctionValue(t *testing.T) {
	value, err := bcv.GetDollarValue()
	if err != nil {
		t.Error(err)
	}
	if !utils.HasEightDecimals(value) {
		t.Errorf("The value does not have eight decimal places: %s", value)
	}
}

func TestCurrencyValueAsFloat64AmericanDollar(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat64(bcv.AmericanDollar)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat64Euro(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat64(bcv.Euro)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat64TurkishLira(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat64(bcv.TurkishLira)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat64Yuan(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat64(bcv.Yuan)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat64Ruble(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat64(bcv.Ruble)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat32AmericanDollar(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat32(bcv.AmericanDollar)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat32Euro(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat32(bcv.Euro)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat32TurkishLira(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat32(bcv.TurkishLira)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat32Yuan(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat32(bcv.Yuan)
	if err != nil {
		t.Error(err)
	}
}

func TestCurrencyValueAsFloat32Ruble(t *testing.T) {
	_, err := bcv.GetCurrencyValueAsFloat32(bcv.Ruble)
	if err != nil {
		t.Error(err)
	}
}
