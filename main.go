package main

import "fmt"

func main() {
	print("This is a Go function")

	learnDefer()
	UpdateWeddingDetails()
	RunCalculator()
}

func learnDefer() {
	var age int32 = 22

	defer fmt.Println("this is a deferred statement")
	fmt.Println(age, "\n")
}
