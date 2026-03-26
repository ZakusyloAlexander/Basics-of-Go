package main

import "fmt"

func main() {
	// Усі цілочисельні змінні без явної ініціалізації мають значення 0.
	var defaultInt8 int8
	var defaultInt16 int16
	var defaultInt32 int32
	var defaultInt64 int64
	var defaultInt int

	// Виводимо значення за замовчуванням для знакових типів.
	fmt.Println("Default values (signed): ", defaultInt8, defaultInt16, defaultInt32, defaultInt64, defaultInt)

	var defaultuInt8 uint8
	var defaultuInt16 uint16
	var defaultuInt32 uint32
	var defaultuInt64 uint64
	var defaultuInt uint

	// Виводимо значення за замовчуванням для беззнакових типів.
	fmt.Println("Default values (unsigned): ", defaultuInt8, defaultuInt16, defaultuInt32, defaultuInt64, defaultuInt)

	// Завдання: створити цілочисельну змінну (без виведення).
	var myInt int = 42
	// Щоб компілятор не видав помилку "declared and not used", присвоюємо в blank identifier.
	_ = myInt
}
