package main

import (
	"fmt"

	"osifo.dev/learn-go/fileutils"
	"osifo.dev/learn-go/structs"
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
	device001 := structs.Device{Id: "001", Name: "test_device_1"}
	device002 := structs.NewDevice("002", "test_device_2", structs.Coordinate{"lat": 3.234, "long": 4.902})

	fmt.Printf("Device data: %v", device001.GetReadings())
	fmt.Printf("Device data: %v", device002.GetReadings())

}

func main() {
	print("This is a Go function")
	learnDefer()
	UpdateWeddingDetails()
	// RunCalculator()
	readFileData()
	writeDateToFile()
	getDeviceReadings()
}
