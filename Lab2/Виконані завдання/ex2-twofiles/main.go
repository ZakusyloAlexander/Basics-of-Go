// Пункт 4: додаток з двох вихідних файлів у пакеті main.
// Логіка функцій — у funcs.go; тут лише точка входу (функція main).
package main

import "fmt"

func main() {
	// Три дійсні числа для демонстрації Min3 та Average3.
	x, y, z := 5.0, 2.0, 8.0

	fmt.Printf("Числа: %.2f, %.2f, %.2f\n", x, y, z)
	// Min3, Average3 оголошені в funcs.go того ж пакету main — імпорт не потрібен.
	fmt.Printf("Мінімальне значення: %.2f\n", Min3(x, y, z))
	fmt.Printf("Середнє значення: %.2f\n", Average3(x, y, z))

	// Розв'язок рівняння coefA·x + coefB = 0.
	coefA, coefB := 2.0, 4.0
	solution, err := SolveLinear(coefA, coefB)
	if err != nil {
		fmt.Printf("Рівняння %.2f*x + %.2f = 0: %v\n", coefA, coefB, err)
	} else {
		fmt.Printf("Рівняння %.2f*x + %.2f = 0 => x = %.2f\n", coefA, coefB, solution)
	}
}
