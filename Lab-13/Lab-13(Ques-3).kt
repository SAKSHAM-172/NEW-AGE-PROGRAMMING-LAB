class Student(
    private var rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {
    // Accepts/updates student data from user input
    fun acceptData() {
        print("Enter roll number: ")
        rollNo = readln().trim().toInt()

        print("Enter name: ")
        name = readln().trim()

        print("Enter course: ")
        course = readln().trim()

        print("Enter percentage: ")
        percentage = readln().trim().toDouble()
    }

    // Displays all details in a formatted manner
    fun displayDetails() {
        println("-----------------------------")
        println("Roll No    : $rollNo")
        println("Name       : $name")
        println("Course     : $course")
        println("Percentage : ${"%.2f".format(percentage)}%")
        println("-----------------------------")
    }
}

fun main() {
    // Student 1: created with constructor values, displayed, then updated from input
    val s1 = Student(101, "Prabhat", "B.Tech CSE", 78.5)
    println("Student 1 (initial data):")
    s1.displayDetails()

    println("\nUpdate Student 1:")
    s1.acceptData()
    println("\nStudent 1 (updated data):")
    s1.displayDetails()

    // Student 2
    val s2 = Student(102, "Riya", "B.Tech IT", 85.25)
    println("\nStudent 2 (initial data):")
    s2.displayDetails()

    println("\nUpdate Student 2:")
    s2.acceptData()
    println("\nStudent 2 (updated data):")
    s2.displayDetails()
}