fun main() {
    // Task 1
    println("Hello Kotlin!!!")

    // Task 2: Employee bonus
    print("Enter department code (IT/HR/FIN): ")
    val dept = readln().trim().uppercase()
    print("Enter years of experience: ")
    val years = readln().toInt()
    print("Enter salary: ")
    val salary = readln().toDouble()

    val rate = when {
        dept == "IT" && years > 5 -> 20
        dept == "HR" && years > 3 -> 15
        dept == "FIN" && years > 4 -> 12
        else -> 10
    }
    val bonus = salary * rate / 100
    println("Bonus rate: $rate%, Bonus amount: $bonus")

    // Task 3: Average of digits
    print("Enter an integer: ")
    val n = kotlin.math.abs(readln().toInt())
    val digits = n.toString().map { it.digitToInt() }
    println("Average of digits: ${digits.sum().toDouble() / digits.size}")
}