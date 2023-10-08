package dtwin

import (
	"encoding/json"
	"fmt"
	"time"
)

type HealthStatus struct {
	Device
	Timestamp time.Time `json:"timestamp"`
	Metadata  string    `json:"metadata"`
	IsOK      bool      `json:"isOk"`
	Status    string    `json:"status"`
}

func (status HealthStatus) GetData() string {
	return fmt.Sprintf("%v", status)
}

func (status HealthStatus) String() string {
	val, err := json.MarshalIndent(status, "", "  ")

	if err != nil {
		return fmt.Sprintf("An error occured while converting this reading to a string:\n%v", err)
	}
	return string(val)
}
