class Person(val name: String, val email: String?)

fun main() {
    // ---------- (1) Nullable vs non-nullable ----------
    println("===== 1. Nullable vs Non-Nullable =====")

    var nonNullable: String = "Kotlin"
    var nullable: String? = "Kotlin"

    println("nonNullable = $nonNullable")
    println("nullable    = $nullable")

    nullable = null                 // allowed: type is String?
    println("nullable after assigning null = $nullable")

    // nonNullable = null           // COMPILE ERROR: Null can not be a value of a non-null type String
    // println(nullable.length)     // COMPILE ERROR: only safe (?.) or non-null asserted (!!.) calls are allowed

    println("nonNullable.length = ${nonNullable.length}")   // direct access is fine

    // ---------- (2) Safe call operator ?. ----------
    println("\n===== 2. Safe Call Operator (?.) =====")

    val p1 = Person("Aarav", "aarav@example.com")
    val p2 = Person("Diya", null)

    println("${p1.name}'s email length: ${p1.email?.length}")   // 17
    println("${p2.name}'s email length: ${p2.email?.length}")   // null, no crash

    println("${p1.name}'s email uppercase: ${p1.email?.uppercase()}")
    println("${p2.name}'s email uppercase: ${p2.email?.uppercase()}")

    // ---------- (3) Elvis operator ?: ----------
    println("\n===== 3. Elvis Operator (?:) =====")

    val email1 = p1.email ?: "no email provided"
    val email2 = p2.email ?: "no email provided"
    println("${p1.name}: $email1")
    println("${p2.name}: $email2")

    val length1 = p1.email?.length ?: 0
    val length2 = p2.email?.length ?: 0
    println("Length with default -> ${p1.name}: $length1, ${p2.name}: $length2")

    // ---------- (4) Not-null assertion !! ----------
    println("\n===== 4. Not-Null Assertion (!!) =====")

    val safeValue: String? = "Hello"
    println("Force-unwrapping non-null value: ${safeValue!!.length}")

    val nullValue: String? = null
    try {
        println("Force-unwrapping null value...")
        println(nullValue!!.length)
    } catch (e: NullPointerException) {
        println("Caught NullPointerException: ${e.message}")
    }
}