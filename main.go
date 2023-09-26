package main

import (
	"fmt"

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

func main() {
	print("This is a Go function")
	learnDefer()
	UpdateWeddingDetails()
	// RunCalculator()
	readFileData()
	writeDateToFile()
}
