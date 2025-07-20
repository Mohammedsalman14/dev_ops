package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: <number1> <operation> <number2>")
		fmt.Println("Example: 10 + 5")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	op := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers")
		return
	}

	switch op {
	case "+":
		fmt.Printf("%.2f\n", num1+num2)
	case "-":
		fmt.Printf("%.2f\n", num1-num2)
	case "*":
		fmt.Printf("%.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Printf("%.2f\n", num1/num2)
	default:
		fmt.Println("Invalid operation. Use +, -, *, or /")
	}
}
