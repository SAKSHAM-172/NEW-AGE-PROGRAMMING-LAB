package main

import (
	"fmt"
	"math"
	"strings"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by zero", a)
	}
	return a / b, nil
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot take square root of negative number %.2f", x)
	}
	return math.Sqrt(x), nil
}

func Greet(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

func main() {
	if result, err := Divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide(10, 2) =", result)
	}

	if result, err := Divide(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide(10, 0) =", result)
	}

	if result, err := SquareRoot(16); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("SquareRoot(16) =", result)
	}

	if result, err := SquareRoot(-9); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("SquareRoot(-9) =", result)
	}

	if msg, err := Greet("Asha"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(msg)
	}

	if msg, err := Greet("   "); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(msg)
	}
}
