// Reverse Polish Notation calculator
// Usage: rpncalc <expression>
// Example: rpncalc "2 3 + 4 *"
// Output: 20
// Example: rpncalc "2 3 4 + *"
// Output: 14
//
// Author: Michael John
// Email: michael @ aka.ninja

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check we have an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: rpncalc <expression>")
		os.Exit(1)
	}

	// evaluate the expression
	result, err := evalRPN(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", result)
}

// Evaluate the expression
// Returns the result of the expression, or an error if the expression is invalid
func evalRPN(expr string) (float64, error) {
	tokens := strings.Fields(expr)
	var stack Stack

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if stack.Len() < 2 {
				return 0, fmt.Errorf("not enough values for the operation '%v'", token)
			}

			// Pop the top two elements from the stack
			// we can ignore the second return value as we know the stack is not empty
			b, _ := stack.Pop()
			a, _ := stack.Pop()

			// Perform the operation
			result, err := performOperation(a, b, token)
			if err != nil {
				return 0, err
			}

			// Push the result back onto the stack
			stack.Push(result)
		default:
			// Convert the token to a float64 and push it onto the stack
			val, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token '%v'", token)
			}
			stack.Push(val)
		}
	}

	// The result should be the only element on the stack
	if stack.Len() != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	// Pop the result from the stack
	result, _ := stack.Pop()
	return result, nil
}

func performOperation(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation '%v'", op)
	}
}

