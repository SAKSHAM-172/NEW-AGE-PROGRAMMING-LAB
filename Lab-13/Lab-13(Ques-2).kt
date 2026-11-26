// (1) Default parameters
fun greet(name: String = "Student", greeting: String = "Hello", punctuation: String = "!") {
    println("$greeting, $name$punctuation")
}

// (2) Named arguments: reuses a function with several parameters
fun createProfile(name: String, age: Int, city: String) {
    println("Name: $name, Age: $age, City: $city")
}

// (3) Vararg
fun printNumbers(vararg numbers: Int) {
    println("Received ${numbers.size} numbers:")
    for (n in numbers) {
        print("$n ")
    }
    println()
}

// (4) Higher-order function
fun applyToList(list: List<Int>, operation: (Int) -> Int): List<Int> {
    return list.map(operation)
}

fun main() {
    println("--- 1. Default parameters ---")
    greet()                              // uses all defaults
    greet("Prabhat")                     // overrides name only
    greet("Prabhat", "Welcome", "!!!")   // overrides all

    println("\n--- 2. Named arguments ---")
    createProfile(name = "Prabhat", age = 20, city = "Ghaziabad")
    createProfile(city = "Delhi", name = "Amit", age = 21)   // out of order
    greet(punctuation = "?", name = "Riya")                  // named + defaults

    println("\n--- 3. Vararg ---")
    printNumbers(1, 2, 3)
    printNumbers(10, 20, 30, 40, 50)
    printNumbers()                       // zero arguments is also valid

    println("\n--- 4. Higher-order function ---")
    val nums = listOf(1, 2, 3, 4, 5)
    println("Original: $nums")
    println("Squared: ${applyToList(nums) { it * it }}")
    println("Doubled: ${applyToList(nums) { x -> x * 2 }}")
    val addTen: (Int) -> Int = { it + 10 }
    println("Plus ten: ${applyToList(nums, addTen)}")
}