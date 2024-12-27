package bcv

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/shopspring/decimal"
)

const (
	Euro int8 = iota
	Yuan
	TurkishLira
	Ruble
	AmericanDollar
)

func GetCurrencyValue(currency int8) (string, error) {
	htmlRequest, err := http.Get("https://www.bcv.org.ve/")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error (go to https://www.bcv.org.ve/): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	doc, docErr := goquery.NewDocumentFromReader(htmlRequest.Body)
	if docErr != nil {
		return "", errors.New(fmt.Sprintf("Error (html.Parse(htmlRequest.Body)): %v", docErr))
	}
	var currencyValue string = "dolar"
	switch currency {
	case Euro:
		currencyValue = "euro"
		break
	case Yuan:
		currencyValue = "yuan"
		break
	case TurkishLira:
		currencyValue = "lira"
		break
	case Ruble:
		currencyValue = "rublo"
		break
	case AmericanDollar:
		break
	default:
		break
	}
	return strings.ReplaceAll(strings.TrimSpace(doc.Find(fmt.Sprintf("div #%s div.col-sm-6.col-xs-6.centrado strong", currencyValue)).Text()), ",", "."), nil
}

func GetDollarValue() (string, error) {
	htmlRequest, err := http.Get("https://www.bcv.org.ve/")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error (go to https://www.bcv.org.ve/): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	doc, docErr := goquery.NewDocumentFromReader(htmlRequest.Body)
	if docErr != nil {
		return "", errors.New(fmt.Sprintf("Error (html.Parse(htmlRequest.Body)): %v", docErr))
	}
	return strings.ReplaceAll(strings.TrimSpace(doc.Find("div #dolar div.col-sm-6.col-xs-6.centrado strong").Text()), ",", "."), nil
}

func GetDollarValueAsFloat64() (float64, error) {
	htmlRequest, err := http.Get("https://www.bcv.org.ve/")
	if err != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (go to https://www.bcv.org.ve/): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.00000000, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	doc, docErr := goquery.NewDocumentFromReader(htmlRequest.Body)
	if docErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (html.Parse(htmlRequest.Body)): %v", docErr))
	}
	dollarValueAsString := strings.ReplaceAll(strings.TrimSpace(doc.Find("div #dolar div.col-sm-6.col-xs-6.centrado strong").Text()), ",", ".")
	dollarValueDecimal, decimalErr := decimal.NewFromString(dollarValueAsString)
	if decimalErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (decimal.NewFromString): %v", decimalErr))
	}
	numberValue, _ := dollarValueDecimal.BigFloat().Float64()
	return numberValue, nil
}

func GetDollarValueAsFloat32() (float32, error) {
	htmlRequest, err := http.Get("https://www.bcv.org.ve/")
	if err != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (go to https://www.bcv.org.ve/): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.00000000, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	doc, docErr := goquery.NewDocumentFromReader(htmlRequest.Body)
	if docErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (html.Parse(htmlRequest.Body)): %v", docErr))
	}
	dollarValueAsString := strings.ReplaceAll(strings.TrimSpace(doc.Find("div #dolar div.col-sm-6.col-xs-6.centrado strong").Text()), ",", ".")
	dollarValueDecimal, decimalErr := decimal.NewFromString(dollarValueAsString)
	if decimalErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (decimal.NewFromString): %v", decimalErr))
	}
	numberValue, _ := dollarValueDecimal.BigFloat().Float32()
	return numberValue, nil
}

func GetCurrencyValueAsFloat64(currency int8) (float64, error) {
	htmlRequest, err := http.Get("https://www.bcv.org.ve/")
	if err != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (go to https://www.bcv.org.ve/): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.00000000, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	doc, docErr := goquery.NewDocumentFromReader(htmlRequest.Body)
	if docErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (html.Parse(htmlRequest.Body)): %v", docErr))
	}
	var currencyValue string = "dolar"
	switch currency {
	case Euro:
		currencyValue = "euro"
		break
	case Yuan:
		currencyValue = "yuan"
		break
	case TurkishLira:
		currencyValue = "lira"
		break
	case Ruble:
		currencyValue = "rublo"
		break
	case AmericanDollar:
		break
	default:
		break
	}
	currencyValueString := strings.ReplaceAll(strings.TrimSpace(doc.Find(fmt.Sprintf("div #%s div.col-sm-6.col-xs-6.centrado strong", currencyValue)).Text()), ",", ".")
	number, decimalErr := decimal.NewFromString(currencyValueString)
	if decimalErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (decimal.NewFromString): %v", decimalErr))
	}
	decimalNumber, _ := number.BigFloat().Float64()
	return decimalNumber, nil
}

func GetCurrencyValueAsFloat32(currency int8) (float32, error) {
	htmlRequest, err := http.Get("https://www.bcv.org.ve/")
	if err != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (go to https://www.bcv.org.ve/): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.00000000, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	doc, docErr := goquery.NewDocumentFromReader(htmlRequest.Body)
	if docErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (html.Parse(htmlRequest.Body)): %v", docErr))
	}
	var currencyValue string = "dolar"
	switch currency {
	case Euro:
		currencyValue = "euro"
		break
	case Yuan:
		currencyValue = "yuan"
		break
	case TurkishLira:
		currencyValue = "lira"
		break
	case Ruble:
		currencyValue = "rublo"
		break
	case AmericanDollar:
		break
	default:
		break
	}
	currencyValueString := strings.ReplaceAll(strings.TrimSpace(doc.Find(fmt.Sprintf("div #%s div.col-sm-6.col-xs-6.centrado strong", currencyValue)).Text()), ",", ".")
	number, decimalErr := decimal.NewFromString(currencyValueString)
	if decimalErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (decimal.NewFromString): %v", decimalErr))
	}
	decimalNumber, _ := number.BigFloat().Float32()
	return decimalNumber, nil
}
