package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func (p *Person) ReadData() {
	p.Name = readLine("Enter name: ")

	for {
		age, err := strconv.Atoi(readLine("Enter age: "))
		if err != nil || age <= 0 {
			fmt.Println("Invalid age, try again.")
			continue
		}
		p.Age = age
		break
	}

	p.Job = readLine("Enter job: ")

	for {
		salary, err := strconv.ParseFloat(readLine("Enter salary: "), 64)
		if err != nil || salary < 0 {
			fmt.Println("Invalid salary, try again.")
			continue
		}
		p.Salary = salary
		break
	}
}

func (p Person) Display() {
	fmt.Println("-------------------------")
	fmt.Printf("Name   : %s\n", p.Name)
	fmt.Printf("Age    : %d\n", p.Age)
	fmt.Printf("Job    : %s\n", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
	fmt.Println("-------------------------")
}

func main() {
	var p1, p2 Person

	fmt.Println("Enter details for Person 1")
	p1.ReadData()

	fmt.Println("\nEnter details for Person 2")
	p2.ReadData()

	fmt.Println("\nPerson 1 details:")
	p1.Display()

	fmt.Println("\nPerson 2 details:")
	p2.Display()
}
