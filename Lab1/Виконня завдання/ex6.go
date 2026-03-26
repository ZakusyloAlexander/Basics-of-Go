package main

import "fmt"

func main() {
	// Використовуємо uint8, щоб наочно бачити 8-бітні шаблони.
	var x, y, z uint8

	// Початкові значення для прикладів.
	x = 9
	y = 28
	z = x // Зберігаємо оригінал x для повторних операцій.

	fmt.Println("Бітові операції")

	// Виводимо результат і в десятковому, і в двійковому вигляді.
	fmt.Printf("^x      = ^(%d)      = ^(%.8b)            = %.8b = %d\n", x, x, ^x, ^x)
	fmt.Printf("x << 2  = (%d << 2)  = (%.8b << 2)        = %.8b = %d\n", x, x, x<<2, x<<2)
	fmt.Printf("x >> 2  = (%d >> 2)  = (%.8b >> 2)        = %.8b = %d\n", x, x, x>>2, x>>2)
	fmt.Printf("x & y   = (%d & %d)  = (%.8b & %.8b)  = %.8b = %d\n", x, y, x, y, x&y, x&y)
	fmt.Printf("x | y   = (%d | %d)  = (%.8b | %.8b)  = %.8b = %d\n", x, y, x, y, x|y, x|y)
	fmt.Printf("x ^ y   = (%d ^ %d)  = (%.8b ^ %.8b)  = %.8b = %d\n", x, y, x, y, x^y, x^y)
	fmt.Printf("x &^ y  = (%d &^ %d) = (%.8b &^ %.8b) = %.8b = %d\n", x, y, x, y, x&^y, x&^y)
	fmt.Printf("x %% y   = (%d %% %d)  = (%.8b %% %.8b)  = %.8b = %d\n", x, y, x, y, x%y, x%y)

	fmt.Println("\nБітові операції з присвоюванням")

	// Перед кожною новою операцією відновлюємо x у вихідне значення.
	x = z
	x &= y
	fmt.Printf("x &= y   = (%d &= %d)  = (%.8b &= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x |= y
	fmt.Printf("x |= y   = (%d |= %d)  = (%.8b |= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x ^= y
	fmt.Printf("x ^= y   = (%d ^= %d)  = (%.8b ^= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x &^= y
	fmt.Printf("x &^= y  = (%d &^= %d) = (%.8b &^= %.8b) = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x %= y
	fmt.Printf("x %%= y   = (%d %%= %d)  = (%.8b %%= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)

	// Завдання: коротко пояснити результати операцій.
	fmt.Println("\nПояснення операцій:")
	fmt.Println("^x      - бітове НЕ (інверсія всіх бітів)")
	fmt.Println("x << 2  - зсув вліво на 2 біти (множення на 2^2 = 4)")
	fmt.Println("x >> 2  - зсув вправо на 2 біти (ділення на 2^2 = 4)")
	fmt.Println("x & y   - бітове І (AND) - 1 тільки якщо обидва біти = 1")
	fmt.Println("x | y   - бітове АБО (OR) - 1 якщо хоча б один біт = 1")
	fmt.Println("x ^ y   - бітове виключне АБО (XOR) - 1 якщо біти різні")
	fmt.Println("x &^ y  - бітове очищення (AND NOT) - очищає біти x де y=1")
	fmt.Println("x % y   - залишок від ділення")
}
