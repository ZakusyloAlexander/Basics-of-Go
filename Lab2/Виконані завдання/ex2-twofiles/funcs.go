// Пункт 4: три функції в окремому файлі того ж пакету main.
// Усі .go файли з package main в одній директорії збираються в одну програму.
package main

import "errors"

// Min3 повертає мінімальне з трьох чисел.
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

// Average3 — середнє арифметичне трьох чисел.
func Average3(a, b, c float64) float64 {
	return (a + b + c) / 3
}

// SolveLinear розв'язує ax + b = 0; при a = 0 повертає помилку через errors.New.
func SolveLinear(a, b float64) (float64, error) {
	if a == 0 {
		if b == 0 {
			return 0, errors.New("рівняння має нескінченно багато розв'язків")
		}
		return 0, errors.New("рівняння не має розв'язку")
	}
	return -b / a, nil
}
