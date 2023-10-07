package cryptoapi_models

import (
	"encoding/json"
	"fmt"
)

type Rate struct {
	Timestamp string  `json:"timestamp"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	Last      string  `json:"last"`
}

func (rate Rate) String() string {
	val, err := json.MarshalIndent(rate, "", "  ")

	if err != nil {
		return fmt.Sprintf("could not marsall rate to json\n%v", err)
	}
	return string(val)
}
