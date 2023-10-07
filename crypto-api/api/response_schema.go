package cryptoapi

import (
	"encoding/json"
	"fmt"
)

type HttpResponse struct {
	Success bool  `json:"success"`
	Data    any   `json:"data"`
	Error   error `json:"error"`
}

func (response HttpResponse) String() string {
	value, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		return fmt.Sprintf("could not marshal http response to json:%v\n", err)
	}

	return string(value)
}
