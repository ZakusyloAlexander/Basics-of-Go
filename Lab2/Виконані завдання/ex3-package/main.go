// Пункт 5: головний модуль (package main) підключає окремий пакет calc.
// Ім'я модуля задається в go.mod; шлях імпорту має збігатися з module/підпакет.
package main

import (
	"fmt"

	"lab2ex3/calc"
)

func main() {
	x, y, z := 5.0, 2.0, 8.0

	fmt.Printf("Числа: %.2f, %.2f, %.2f\n", x, y, z)
	// Функції з пакету calc викликаються з префіксом calc.
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
