// Пункт 5–6: пакет calc з трьома функціями; для нього написані unit-тести.
package calc

import "errors"

// Min3 повертає мінімальне значення з трьох елементів.
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

// Average3 обчислює середнє значення трьох елементів.
func Average3(a, b, c float64) float64 {
	return (a + b + c) / 3
}

// SolveLinear розв'язує рівняння першого порядку ax + b = 0.
func SolveLinear(a, b float64) (float64, error) {
	if a == 0 {
		if b == 0 {
			return 0, errors.New("рівняння має нескінченно багато розв'язків")
		}
		return 0, errors.New("рівняння не має розв'язку")
	}
	return -b / a, nil
}
