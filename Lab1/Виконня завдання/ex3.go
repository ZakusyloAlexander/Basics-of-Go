package main

import "fmt"

func main() {
	// Оголошуємо змінні різних цілочисельних типів.
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	// Тип userautoinit буде визначений автоматично (тут int).
	var userautoinit = -4

	// Виводимо значення змінних.
	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення (:=) працює лише для нових змінних.
	intVar := 10

	// %d - десяткове значення, %T - тип змінної.
	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Завдання 1: вивести типи всіх змінних.
	fmt.Printf("userinit8 Type = %T\n", userinit8)
	fmt.Printf("userinit16 Type = %T\n", userinit16)
	fmt.Printf("userinit64 Type = %T\n", userinit64)
	fmt.Printf("userautoinit Type = %T\n", userautoinit)
	fmt.Printf("intVar Type = %T\n\n", intVar)

	// Завдання 2: присвоюємо intVar значення інших типів із явним приведенням.
	intVar = int(userinit16)
	fmt.Printf("intVar = userinit16: Value = %d Type = %T\n", intVar, intVar)
	
	intVar = int(userautoinit)
	fmt.Printf("intVar = userautoinit: Value = %d Type = %T\n", intVar, intVar)
}
