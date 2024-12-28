package bcv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jeremiasbots/currencyRatesLib/types"
	"github.com/shopspring/decimal"
)

const (
	Euro int8 = iota
	Yuan
	TurkishLira
	Ruble
	AmericanDollar
)

type BCVApiResponse types.APIResponse

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
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/oficial")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error (go to https://ve.dolarapi.com/v1/dolares/oficial): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return "", errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *BCVApiResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return "", errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	roundedNumber := math.Round(dataResponse.Average*1e8) / 1e8
	return fmt.Sprintf("%.8f", roundedNumber), nil
}

func GetDollarData() (BCVApiResponse, error) {
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/oficial")
	if err != nil {
		return BCVApiResponse{}, errors.New(fmt.Sprintf("Error (go to https://ve.dolarapi.com/v1/dolares/oficial): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return BCVApiResponse{}, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return BCVApiResponse{}, errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *BCVApiResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return BCVApiResponse{}, errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	return *dataResponse, nil
}

func GetDollarValueAsFloat64() (float64, error) {
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/oficial")
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Error (go to https://ve.dolarapi.com/v1/dolares/oficial): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.0, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return 0.0, errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *BCVApiResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return 0.0, errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	return dataResponse.Average, nil
}

func GetDollarValueAsFloat32() (float32, error) {
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/oficial")
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Error (go to https://ve.dolarapi.com/v1/dolares/oficial): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.0, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return 0.0, errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *BCVApiResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return 0.0, errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	averageResponse, _ := decimal.NewFromFloat(dataResponse.Average).BigFloat().Float32()
	return averageResponse, nil
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
