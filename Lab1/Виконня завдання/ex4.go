package main

import "fmt"

func main() {
	// Початкові значення для арифметичних операцій.
	var a int16 = 2
	var b int16 = 3
	var c int64 = 10

	// Базові арифметичні операції.
	fmt.Println("a + b        = ", a+b)
	fmt.Println("a - b        = ", a-b)
	fmt.Println("a * b        = ", a*b)
	// Для int-ділення результат дробової частини відкидається.
	fmt.Println("int(a / b)   = ", int(a/b), "\n")

	// Демонстрація скорочених операторів зміни значення.
	c--
	fmt.Println("c--     = ", c)
	c++
	fmt.Println("c++     = ", c)
	c += 10
	fmt.Println("c += 10 = ", c)
	c -= 5
	fmt.Println("c -= 5  = ", c)
	c *= 3
	fmt.Println("c *= 3  = ", c)
	c /= 7
	fmt.Println("c /= 7  = ", c)

	// Завдання: показати c--.
	// В Go ++/-- не можна використовувати всередині виразу, лише окремим оператором.
	c--
	fmt.Println("c--     = ", c)
}
