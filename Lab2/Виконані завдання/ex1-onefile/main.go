// Пункт 3: програма в одному вихідному файлі з трьома функціями.
// Усі функції та main знаходяться в пакеті main — це точка входу програми.
package main

import (
	"errors"
	"fmt"
)

// Min3 повертає найменше з трьох дійсних чисел.
// Параметри a, b, c мають один тип float64 (скорочений запис списку параметрів).
func Min3(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

// Average3 обчислює середнє арифметичне трьох чисел: (a+b+c)/3.
func Average3(a, b, c float64) float64 {
	return (a + b + c) / 3
}

// SolveLinear розв'язує лінійне рівняння ax + b = 0 відносно x.
// Повертає два значення: x та error.
// Якщо a ≠ 0, то x = -b/a; якщо a = 0 — повертаємо помилку (немає розв'язку або їх безліч).
func SolveLinear(a, b float64) (float64, error) {
	if a == 0 {
		if b == 0 {
			// 0·x + 0 = 0 — тотожність, розв'язків нескінченно багато.
			return 0, errors.New("рівняння 0*x + 0 = 0 має нескінченно багато розв'язків")
		}
		// 0·x + b = 0 при b ≠ 0 неможливо.
		return 0, errors.New("рівняння 0*x + b = 0 не має розв'язку")
	}
	return -b / a, nil
}

func main() {
	// Тестові числа для Min3 та Average3.
	x, y, z := 5.0, 2.0, 8.0

	fmt.Printf("Числа: %.2f, %.2f, %.2f\n", x, y, z)
	fmt.Printf("Мінімальне значення: %.2f\n", Min3(x, y, z))
	fmt.Printf("Середнє значення: %.2f\n", Average3(x, y, z))

	// Коефіцієнти лінійного рівняння: 2x + 4 = 0 => x = -2.
	a, b := 2.0, 4.0
	solution, err := SolveLinear(a, b)
	if err != nil {
		// Якщо err != nil, розв'язку немає або він неоднозначний.
		fmt.Printf("Рівняння %.2f*x + %.2f = 0: %v\n", a, b, err)
	} else {
		fmt.Printf("Рівняння %.2f*x + %.2f = 0 => x = %.2f\n", a, b, solution)
	}

	// Випадок a = 0: рівняння 0·x + 5 = 0 не має розв'язку.
	solution2, err2 := SolveLinear(0, 5)
	if err2 != nil {
		fmt.Printf("Рівняння 0*x + 5 = 0: %v\n", err2)
	} else {
		fmt.Printf("x = %.2f\n", solution2)
	}
}
