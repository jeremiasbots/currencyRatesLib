package bcv

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
