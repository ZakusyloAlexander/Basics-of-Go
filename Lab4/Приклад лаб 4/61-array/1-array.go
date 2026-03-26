package main

import (
	"fmt"
)

func main() {
	// Определение переменных
	var first [2]bool
	second := [...]int{-1, -2, -3}
	third := [6]uint{1, 2, 3, 4, 5, 6}
	var fourth [2][2]float32
	fifth := [2][2]float32{{1}, {5.4, 7.7}}
	//fifth := [2][2]float32{1, 0, 5.4, 7.7}

	fmt.Printf("first  len(%d) = %v\n", len(first), first)
	fmt.Printf("second len(%d) = %v\n", len(second), second)
	fmt.Printf("third  len(%d) = %v\n", len(third), third)
	fmt.Printf("fourth len(%d) = %v\n", len(fourth), fourth)
	fmt.Printf("fifth  len(%d) = %v\n\n", len(fifth), fifth)

	fmt.Println("second[2]   = ", second[2])
	fmt.Println("fifth[0][1] = ", fifth[0][1])

	copyOfThird := third // Передаются по значению
	third[4] = 500
	third[5] = 600
	fmt.Println("\nthird       = ", third)
	fmt.Println("copyOfThird = ", copyOfThird)

	ptrThird := &third // Указатель на массив
	ptrThird[4] = 5
	ptrThird[5] = 6
	fmt.Println("\nthird       = ", third)
	fmt.Println("ptrThird = ", *ptrThird)
}
