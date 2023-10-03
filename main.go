package main

import (
	"fmt"
	"time"

	"osifo.dev/learn-go/dtwin"
	"osifo.dev/learn-go/fileutils"
)

func readFileData() {
	content, err := fileutils.ReadTextfile("fileutils/data/user.json")

	if err != nil {
		fmt.Printf("An error occurred while reading file: \n %s", err)
	} else {
		fmt.Printf("\n %s \n", content)
	}
}

func writeDateToFile() {
	error := fileutils.WriteToFile("fileutils/data/output.txt", "this is a test output")

	if error != nil {
		fmt.Printf("Something went wrong while writing data to specified file:\n%s", error)
	} else {
		fmt.Println("Data written to file succesfully")
	}
}

func learnDefer() {
	var age int32 = 22

	defer fmt.Println("this is a deferred statement")
	fmt.Println(age)
}

func getDeviceReadings() {
	device001 := dtwin.Device{Id: "001", Name: "test_device_1"}
	device002 := dtwin.NewDevice("002", "test_device_2", dtwin.Coordinate{"lat": 3.234, "long": 4.902})

	fmt.Printf("Device data: %v", device001.GetReadings())
	fmt.Printf("Device data: %v", device002.GetReadings())

}

func DisplayDataFromDevice() {
	device001 := dtwin.Device{Id: "001", Name: "test_device_1"}
	device002 := dtwin.NewDevice("002", "test_device_2", dtwin.Coordinate{"lat": 3.234, "long": 4.902})

	reading1 := dtwin.Reading{Device: device001, Identifier: "001", Value: 150, Timestamp: time.Now()}
	reading2 := dtwin.Reading{Device: device002, Identifier: "002", Value: 25, Timestamp: time.Now()}

	statusReport1 := dtwin.HealthStatus{Device: device001, Timestamp: time.Now(), IsOK: true, Status: "OK"}
	statusReport2 := dtwin.HealthStatus{Device: device002, Timestamp: time.Now(), IsOK: false, Status: "NOK", Metadata: "al is gut! This is just a test error message"}

	dataItems := [4]dtwin.Loggable{reading1, reading2, statusReport1, statusReport2}

	fmt.Println("Readings from the available devices")

	for _, data := range dataItems {
		fmt.Println(data.GetData())
	}
}

func main() {
	print("This is a Go function")
	learnDefer()
	UpdateWeddingDetails()
	// RunCalculator()
	readFileData()
	writeDateToFile()
	// getDeviceReadings()
	DisplayDataFromDevice()
}
