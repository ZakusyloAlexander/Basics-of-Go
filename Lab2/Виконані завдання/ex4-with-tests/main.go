// Пункт 5–6: main використовує пакет calc; тести лежать у calc/calc_test.go.
package main

import (
	"fmt"

	"lab2ex4/calc"
)

func main() {
	x, y, z := 5.0, 2.0, 8.0

	fmt.Printf("Числа: %.2f, %.2f, %.2f\n", x, y, z)
	fmt.Printf("Мінімальне значення: %.2f\n", calc.Min3(x, y, z))
	fmt.Printf("Середнє значення: %.2f\n", calc.Average3(x, y, z))

	coefA, coefB := 2.0, 4.0
	solution, err := calc.SolveLinear(coefA, coefB)
	if err != nil {
		fmt.Printf("Рівняння %.2f*x + %.2f = 0: %v\n", coefA, coefB, err)
	} else {
		fmt.Printf("Рівняння %.2f*x + %.2f = 0 => x = %.2f\n", coefA, coefB, solution)
	}
}
