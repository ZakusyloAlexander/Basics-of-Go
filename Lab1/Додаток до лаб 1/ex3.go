package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 // Такий варіант ініціалізації також можливий

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення змінної
	// тільки для нових змінних
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Завдання.
	// 1. Вивести типи всіх змінних
	// 2. Присвоїти змінній intVar змінні userinit16 і userautoinit. Результат вивести в консоль.
}
