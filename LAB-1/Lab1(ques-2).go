package main

import "fmt"

func main() {
	a, b := 20, 6
	x, y := 7.5, 2.5

	fmt.Println("Integers:", a, b)
	fmt.Println("Add:", a+b)
	fmt.Println("Subtract:", a-b)
	fmt.Println("Multiply:", a*b)

	fmt.Println("Floats:", x, y)
	fmt.Println("Add:", x+y)
	fmt.Println("Subtract:", x-y)
	fmt.Println("Multiply:", x*y)
}
