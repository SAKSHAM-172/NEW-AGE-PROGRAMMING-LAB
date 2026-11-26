package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func increment(p *int) {
	*p = *p + 10
}

func main() {
	fmt.Println("===== POINTER BASICS =====")
	x := 42
	p := &x
	fmt.Println("Value of x:         ", x)
	fmt.Println("Address of x (&x):  ", &x)
	fmt.Println("Pointer p:          ", p)
	fmt.Println("Value via *p:       ", *p)

	*p = 100
	fmt.Println("After *p = 100, x = ", x)

	fmt.Println("\n===== PASS BY REFERENCE =====")
	num := 5
	fmt.Println("Before function call:", num)
	increment(&num)
	fmt.Println("After function call: ", num)

	fmt.Println("\n===== STRUCT WITH new() =====")
	s := new(Student)
	fmt.Println("Initial struct:", *s)

	s.Name = "Asha"
	s.Age = 20
	s.Marks = 88.5
	fmt.Println("After assignment:", *s)

	s.Marks = 91.0
	(*s).Age = 21
	fmt.Println("After modification:", *s)
	fmt.Println("Name:", s.Name, "Age:", s.Age, "Marks:", s.Marks)
}
