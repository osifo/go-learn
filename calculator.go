package main

import "fmt"

func RunCalculator() {
	var operator string
	var operand1, operand2 int

	fmt.Println("======= Simple Calaculator ========")
	fmt.Println(("What operation do you want to perform"))
	fmt.Scanf("%s", &operator)

	fmt.Println("Enter the first operator")
	fmt.Scanf("%d", &operand1)

	fmt.Println("Enter the second operator")
	fmt.Scanf("%d", &operand2)

	fmt.Println("======= Result =======")
	switch operator {
	case "divide":
		fmt.Println(operand1 / operand2)
	case "add":
		fmt.Println(operand1 + operand2)
	case "multiply":
		fmt.Println(operand1 * operand2)
	case "subtract":
		fmt.Println(operand1 - operand2)
	default:
		fmt.Printf("Invalid operator selected: %s", operator)
	}
}
