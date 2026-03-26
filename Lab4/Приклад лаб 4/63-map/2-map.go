package main

import "fmt"

func main() {
	m := map[float32]string{3.14159: "Число Пи"}
	m[2.71828] = "Основание натурального логарифма"

	fmt.Println(m[3.14159])
	fmt.Println(m[2.71828])
}
