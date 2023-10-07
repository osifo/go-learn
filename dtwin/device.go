package dtwin

import (
	"fmt"
)

type Coordinate = map[string]float32 // here i'm using type alias

type Device struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location Coordinate
}

// I'm creating a factory function to construct new objects from this struct
func NewDevice(id string, name string, location Coordinate) Device {
	return Device{Id: id, Name: name, Location: location}
}

func (device Device) GetReadings() string {
	values := [3]float32{2.0, 3.0, 4.0}

	return fmt.Sprintf("\ndevice: %v\nReadings:%v", device, values)
}
