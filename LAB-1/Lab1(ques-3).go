package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minVal = -1000000.0
	maxVal = 1000000.0
)

var reader = bufio.NewReader(os.Stdin)

// readLine prompts and returns a trimmed line of input.
func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// readFloat loops until the user enters a valid, in-range number.
func readFloat(prompt string) float64 {
	for {
		input := readLine(prompt)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input: please enter a numeric value.")
			continue
		}
		if value < minVal || value > maxVal {
			fmt.Printf("Out of range: enter a value between %.0f and %.0f.\n", minVal, maxVal)
			continue
		}
		return value
	}
}

// readInt loops until the user enters a valid, in-range integer.
func readInt(prompt string) int {
	for {
		input := readLine(prompt)

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input: please enter a whole number.")
			continue
		}
		if float64(value) < minVal || float64(value) > maxVal {
			fmt.Printf("Out of range: enter a value between %.0f and %.0f.\n", minVal, maxVal)
			continue
		}
		return value
	}
}

func integerCalculator() {
	fmt.Println("\n--- Integer mode ---")
	a := readInt("Enter first integer: ")
	b := readInt("Enter second integer: ")

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func floatCalculator() {
	fmt.Println("\n--- Floating-point mode ---")
	x := readFloat("Enter first number: ")
	y := readFloat("Enter second number: ")

	fmt.Printf("%.2f + %.2f = %.2f\n", x, y, x+y)
	fmt.Printf("%.2f - %.2f = %.2f\n", x, y, x-y)
	fmt.Printf("%.2f * %.2f = %.2f\n", x, y, x*y)
}

func main() {
	fmt.Println("=== Simple Calculator ===")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Integer calculation")
		fmt.Println("2. Floating-point calculation")
		fmt.Println("3. Exit")

		choice := readLine("Your choice: ")

		switch choice {
		case "1":
			integerCalculator()
		case "2":
			floatCalculator()
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice: enter 1, 2, or 3.")
		}
	}
}
