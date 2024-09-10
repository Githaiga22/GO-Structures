package main

import (
	"errors"
	"fmt"
)

// Calculator struct with multiple operations as methods
type Calculator struct {
	Add      func(a, b int) int
	Subtract func(a, b int) int
	Multiply func(a, b int) int
	Divide   func(a, b int) (int, error) // Return error if dividing by zero
}

func main() {
	// Initialize the Calculator with functions for different operations
	calculator := Calculator{
		Add: func(a, b int) int {
			return a + b
		},
		Subtract: func(a, b int) int {
			return a - b
		},
		Multiply: func(a, b int) int {
			return a * b
		},
		Divide: func(a, b int) (int, error) {
			if b == 0 {
				return 0, errors.New("cannot divide by zero")
			}
			return a / b, nil
		},
	}

	// Test addition
	resultAdd := calculator.Add(10, 5)
	fmt.Println("Addition result:", resultAdd)

	// Test subtraction
	resultSub := calculator.Subtract(10, 5)
	fmt.Println("Subtraction result:", resultSub)

	// Test multiplication
	resultMul := calculator.Multiply(10, 5)
	fmt.Println("Multiplication result:", resultMul)

	// Test division (with error handling)
	resultDiv, err := calculator.Divide(10, 0) // Division by zero to trigger error
	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Division result:", resultDiv)
	}

	// Another division that works
	resultDiv, err = calculator.Divide(10, 2)
	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Division result:", resultDiv)
	}
}
