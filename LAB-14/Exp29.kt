data class Student(val rollNo: Int, val name: String, val course: String, val percentage: Double)

fun main() {
    // ---------- (1) Mutable vs read-only collections ----------
    println("===== PART 1: Collections =====")

    // List
    val readOnlyList: List<String> = listOf("Apple", "Banana", "Cherry")
    val mutableList: MutableList<String> = mutableListOf("Apple", "Banana", "Cherry")

    println("\n--- List ---")
    println("Read-only list: $readOnlyList")
    println("Contains Banana? ${readOnlyList.contains("Banana")}")
    println("Element at index 1: ${readOnlyList[1]}")   // get
    println("get(2): ${readOnlyList.get(2)}")

    mutableList.add("Mango")
    println("After add(Mango): $mutableList")
    mutableList.remove("Apple")
    println("After remove(Apple): $mutableList")
    println("Mutable list contains Mango? ${mutableList.contains("Mango")}")
    println("Mutable list get(0): ${mutableList.get(0)}")

    // Set
    val readOnlySet: Set<Int> = setOf(1, 2, 3, 3, 2)   // duplicates ignored
    val mutableSet: MutableSet<Int> = mutableSetOf(1, 2, 3)

    println("\n--- Set ---")
    println("Read-only set (duplicates removed): $readOnlySet")
    println("Contains 2? ${readOnlySet.contains(2)}")

    mutableSet.add(4)
    mutableSet.add(2)   // already present, ignored
    println("After add(4), add(2): $mutableSet")
    mutableSet.remove(1)
    println("After remove(1): $mutableSet")
    println("Contains 1? ${mutableSet.contains(1)}")

    // Map
    val readOnlyMap: Map<String, Int> = mapOf("Alice" to 90, "Bob" to 75, "Carol" to 82)
    val mutableMap: MutableMap<String, Int> = mutableMapOf("Alice" to 90, "Bob" to 75)

    println("\n--- Map ---")
    println("Read-only map: $readOnlyMap")
    println("Alice's score: ${readOnlyMap["Alice"]}")
    println("get(Bob): ${readOnlyMap.get("Bob")}")
    println("Contains key Carol? ${readOnlyMap.containsKey("Carol")}")
    println("Contains value 100? ${readOnlyMap.containsValue(100)}")

    mutableMap["Dave"] = 68         // add
    mutableMap.put("Eve", 95)       // add
    println("After adding Dave and Eve: $mutableMap")
    mutableMap.remove("Bob")
    println("After remove(Bob): $mutableMap")
    println("Contains key Bob? ${mutableMap.containsKey("Bob")}")

    // ---------- (2) map, filter, reduce ----------
    println("\n===== PART 2: map, filter, reduce =====")

    val numbers = listOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    println("Original list: $numbers")

    val squares = numbers.map { it * it }
    println("map (square each): $squares")

    val evenSquares = squares.filter { it % 2 == 0 }
    println("filter (keep even): $evenSquares")

    val sum = evenSquares.reduce { acc, n -> acc + n }
    println("reduce (sum): $sum")

    // Step-by-step reduce trace
    println("\nreduce step by step:")
    evenSquares.reduce { acc, n ->
        println("  acc=$acc, next=$n -> ${acc + n}")
        acc + n
    }

    // Strings
    val words = listOf("kotlin", "java", "go", "scala", "clojure")
    println("\nOriginal words: $words")

    val upper = words.map { it.uppercase() }
    println("map (uppercase): $upper")

    val longWords = words.filter { it.length > 4 }
    println("filter (length > 4): $longWords")

    val joined = words.reduce { acc, w -> "$acc-$w" }
    println("reduce (join with '-'): $joined")

    // ---------- (3) sortedBy and groupBy ----------
    println("\n===== PART 3: Sorting and groupBy =====")

    val students = listOf(
        Student(101, "Aarav", "BCA", 82.5),
        Student(102, "Diya", "BTech", 91.0),
        Student(103, "Kabir", "BCA", 67.8),
        Student(104, "Meera", "BTech", 74.2),
        Student(105, "Rohan", "MCA", 88.6),
        Student(106, "Sneha", "MCA", 59.9)
    )

    println("\nOriginal list:")
    students.forEach { println("  $it") }

    val ascending = students.sortedBy { it.percentage }
    println("\nSorted by percentage (ascending):")
    ascending.forEach { println("  ${it.name} - ${it.percentage}") }

    val descending = students.sortedByDescending { it.percentage }
    println("\nSorted by percentage (descending):")
    descending.forEach { println("  ${it.name} - ${it.percentage}") }

    val byCourse = students.groupBy { it.course }
    println("\nGrouped by course:")
    for ((course, list) in byCourse) {
        println("  $course:")
        list.forEach { println("    ${it.name} (${it.percentage}%)") }
    }

    // Bonus: groupBy on a computed key
    val byGrade = students.groupBy {
        when {
            it.percentage >= 85 -> "Distinction"
            it.percentage >= 60 -> "First Class"
            else -> "Pass"
        }
    }
    println("\nGrouped by grade:")
    byGrade.forEach { (grade, list) ->
        println("  $grade: ${list.map { it.name }}")
    }
}