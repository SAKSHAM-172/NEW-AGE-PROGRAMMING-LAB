package main

import "fmt"

func addStudent(s []string, name string) []string {
	return append(s, name)
}

func removeByIndex(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		fmt.Println("Invalid index:", index)
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func updateStudent(s []string, index int, name string) {
	if index < 0 || index >= len(s) {
		fmt.Println("Invalid index:", index)
		return
	}
	s[index] = name
}

func sliceDemo() {
	fmt.Println("===== SLICE OPERATIONS =====")

	students := []string{"Asha", "Ravi", "Meena"}
	fmt.Println("Initial slice:      ", students)

	students = addStudent(students, "Kiran")
	fmt.Println("After add (Kiran):  ", students)

	students = addStudent(students, "Neha")
	fmt.Println("After add (Neha):   ", students)

	students = removeByIndex(students, 1)
	fmt.Println("After remove idx 1: ", students)

	updateStudent(students, 0, "Ananya")
	fmt.Println("After update idx 0: ", students)

	students = removeByIndex(students, 10)
	fmt.Println("Final slice:        ", students)
	fmt.Println("Length:", len(students), "Capacity:", cap(students))
}

func mapDemo() {
	fmt.Println("\n===== MAP OPERATIONS =====")

	marks := make(map[string]int)

	marks["Maths"] = 85
	fmt.Println("After insert Maths:  ", marks)
	marks["Physics"] = 78
	fmt.Println("After insert Physics:", marks)
	marks["Chemistry"] = 92
	fmt.Println("After insert Chem:   ", marks)

	subject := "Physics"
	if value, ok := marks[subject]; ok {
		fmt.Printf("Lookup %s: %d\n", subject, value)
	} else {
		fmt.Printf("Lookup %s: not found\n", subject)
	}

	subject = "History"
	if value, ok := marks[subject]; ok {
		fmt.Printf("Lookup %s: %d\n", subject, value)
	} else {
		fmt.Printf("Lookup %s: not found\n", subject)
	}

	marks["Maths"] = 90
	fmt.Println("After update Maths:  ", marks)

	delete(marks, "Physics")
	fmt.Println("After delete Physics:", marks)

	fmt.Println("Total subjects:", len(marks))
}

func main() {
	sliceDemo()
	mapDemo()
}
