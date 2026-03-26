package main

import "fmt"

func main() {
	// int8 може зберігати ASCII-символи (1 байт).
	var chartype int8 = 'R'

	// %c друкує символ, %d друкує його числовий код.
	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання 1: вивести українську літеру 'Ї'.
	// Для Unicode-символів використовуємо rune (синонім int32).
	var ukrainianLetter rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrainianLetter, ukrainianLetter)
	
	// Повторна демонстрація з rune: int8 не підходить для багатобайтових символів.
	var ukrainianLetterRune rune = 'Ї'
	fmt.Printf("Українська літера '%c' має код: %d\n", ukrainianLetterRune, ukrainianLetterRune)
	
	// Завдання 2: пояснення призначення типу rune.
	fmt.Println("\nПояснення типу 'rune':")
	fmt.Println("rune - це псевдонім для int32 в Go")
	fmt.Println("Призначений для зберігання Unicode код-пойнтів (символів)")
	fmt.Println("Дозволяє коректно працювати з багатобайтовими символами (наприклад, українськими літерами)")
	fmt.Println("На відміну від int8 (byte), який може зберігати тільки 1 байт (0-255),")
	fmt.Println("rune може зберігати будь-який Unicode символ (до 4 байт)")
}
