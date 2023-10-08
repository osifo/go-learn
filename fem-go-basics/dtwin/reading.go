package dtwin

import (
	"encoding/json"
	"fmt"
	"time"
)

type Reading struct {
	Device
	Identifier string    `json:"id"`
	Value      int       `json:"value"`
	Timestamp  time.Time `json:"timestamp"`
}

func (reading Reading) GetData() string {
	return fmt.Sprintf("%v", reading)
}

func (reading Reading) String() string {
	val, err := json.MarshalIndent(reading, "", "  ")

	if err != nil {
		return fmt.Sprintf("An error occured while converting this reading to a string:\n%v", err)
	}
	return string(val)
}
