package main

import "fmt"

func main() {
	// first і second не ініціалізовані явно, тому мають false.
	var first, second bool
	// Явна ініціалізація булевого значення.
	var third bool = true
	// Логічне НЕ інвертує third.
	fourth := !third
	var fifth = true

	// Демонстрація базових логічних значень.
	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	// Унарний оператор заперечення.
	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	// Логічне І (&&).
	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	// Логічне АБО (||).
	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	// Операції порівняння повертають bool.
	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	// Завдання: пояснення отриманих результатів.
	fmt.Println("Пояснення результатів операцій:")
	fmt.Println("\nЛогічні операції:")
	fmt.Println("! - логічне НЕ (інверсія): !true = false, !false = true")
	fmt.Println("&& - логічне І (AND): повертає true тільки якщо обидва операнди true")
	fmt.Println("|| - логічне АБО (OR): повертає true якщо хоча б один операнд true")
	fmt.Println("\nОперації порівняння:")
	fmt.Println("<  - менше")
	fmt.Println(">  - більше")
	fmt.Println("<= - менше або дорівнює")
	fmt.Println(">= - більше або дорівнює")
	fmt.Println("== - дорівнює")
	fmt.Println("!= - не дорівнює")
}
