package monitordollar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/jeremiasbots/currencyRatesLib/go/types"
)

type MonitorDollarAPIResponse types.APIResponse

func GetDollarValue() (string, error) {
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/paralelo")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error (go to https://laslucas.com/?c=m): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return "", errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *MonitorDollarAPIResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return "", errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	roundedNumber := math.Round(dataResponse.Average*1e8) / 1e8
	return fmt.Sprintf("%.8f", roundedNumber), nil
}

func GetDollarValueAsFloat64() (float64, error) {
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/paralelo")
	if err != nil {
		return 0.00000000, errors.New(fmt.Sprintf("Error (go to https://ve.dolarapi.com/v1/dolares/paralelo): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return 0.00000000, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *MonitorDollarAPIResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return 0.00000000, errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	roundedNumber := math.Round(dataResponse.Average*1e8) / 1e8
	return roundedNumber, nil
}

func GetDollarData() (MonitorDollarAPIResponse, error) {
	htmlRequest, err := http.Get("https://ve.dolarapi.com/v1/dolares/paralelo")
	if err != nil {
		return MonitorDollarAPIResponse{}, errors.New(fmt.Sprintf("Error (go to https://ve.dolarapi.com/v1/dolares/paralelo): %v", err))
	}
	defer htmlRequest.Body.Close()
	if htmlRequest.StatusCode != 200 {
		return MonitorDollarAPIResponse{}, errors.New(fmt.Sprintf("htmlRequest.StatusCode error: %d %s", htmlRequest.StatusCode, htmlRequest.Status))
	}
	bodyData, readerErr := io.ReadAll(htmlRequest.Body)
	if readerErr != nil {
		return MonitorDollarAPIResponse{}, errors.New(fmt.Sprintf("io.ReadAll(htmlRequest.Body) error: %v", readerErr))
	}
	var dataResponse *MonitorDollarAPIResponse
	if jsonErr := json.Unmarshal(bodyData, &dataResponse); jsonErr != nil {
		return MonitorDollarAPIResponse{}, errors.New(fmt.Sprintf("json.Unmarshal(bodyData, &dataResponse) error: %v", jsonErr))
	}
	return *dataResponse, nil
}
