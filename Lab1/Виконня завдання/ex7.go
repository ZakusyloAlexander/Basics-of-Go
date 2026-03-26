package main

import "fmt"

func main() {
	// Змінні різних цілих типів.
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	// Демонстрація явного приведення типів і двійкового представлення.
	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання: створити змінні різних типів і виконати арифметичні операції.
	var a int32 = 100
	var b int64 = 50
	
	fmt.Println("\nАрифметичні операції з різними типами:")
	fmt.Printf("a (int32) = %d\n", a)
	fmt.Printf("b (int64) = %d\n", b)
	
	// Для операцій типи мають збігатися, тому a приводимо до int64.
	fmt.Printf("int64(a) + b = %d\n", int64(a)+b)
	fmt.Printf("int64(a) - b = %d\n", int64(a)-b)
	fmt.Printf("int64(a) * b = %d\n", int64(a)*b)
	fmt.Printf("int64(a) / b = %d\n", int64(a)/b)
	fmt.Printf("int64(a) %% b = %d\n", int64(a)%b)
}
