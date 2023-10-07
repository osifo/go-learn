package cryptoapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	cryptoapi_models "osifo.dev/learn-go/crypto-api/models"
)

const apiURL string = "https://cex.io/api/ticker/%v/USD"

func handleError(message string) error {
	return fmt.Errorf("\nError: %v\n", message)
}

func handleJsonDecodeError(err error) error {
	if err != nil {
		errorMessage := fmt.Sprintf("Something went wrong while decoding the json respoonse:%v", err)

		return handleError(errorMessage)
	} else {
		return nil
	}
}

func GetRates(cryptoSymbol string) (response *HttpResponse, err error) {
	url := fmt.Sprintf(apiURL, strings.ToUpper(cryptoSymbol))
	res, err := http.Get(url)

	if err != nil {
		return nil, handleError(fmt.Sprintf("Could not get rates: %v\n", err))
	}

	responseBody, ioError := io.ReadAll(res.Body)
	if ioError != nil {
		return nil, handleError(fmt.Sprintf("Could not read response body: %v\n", ioError))
	}

	if res.StatusCode == http.StatusOK {
		var successResponse cryptoapi_models.Rate
		jsonDecodeError := json.Unmarshal(responseBody, &successResponse)

		err = handleJsonDecodeError(jsonDecodeError)
		if err != nil {
			return nil, handleError(fmt.Sprintf("%v\n", err))
		}

		return &HttpResponse{
				Success: true,
				Data:    successResponse},
			nil

	} else {
		return &HttpResponse{
			Success: false,
			Data:    string(responseBody[:])}, nil
	}
}
