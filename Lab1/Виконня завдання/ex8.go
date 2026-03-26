package main

// Імпорт кількох пакетів.
import (
	"fmt"
	"math"
)

func main() {
	// defaultFloat отримає значення за замовчуванням 0.
	var defaultFloat float32
	// Явна ініціалізація змінної типу float64.
	var defaultDouble float64 = 5.5

	// Вивід значень і типів з плаваючою крапкою.
	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	// Межі float32 та float64 із пакета math.
	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання: змінні через короткий запис і змінні зі значеннями за замовчуванням.
	fmt.Println("Змінні з коротким записом:")
	shortInt := 42
	shortFloat := 3.14
	shortString := "Привіт"
	shortBool := true
	
	fmt.Printf("shortInt (%T) = %d\n", shortInt, shortInt)
	fmt.Printf("shortFloat (%T) = %f\n", shortFloat, shortFloat)
	fmt.Printf("shortString (%T) = %s\n", shortString, shortString)
	fmt.Printf("shortBool (%T) = %t\n\n", shortBool, shortBool)
	
	fmt.Println("Змінні з ініціалізацією за замовчуванням:")
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultString string
	var defaultBool bool
	
	fmt.Printf("defaultInt (%T) = %d\n", defaultInt, defaultInt)
	fmt.Printf("defaultFloat32 (%T) = %f\n", defaultFloat32, defaultFloat32)
	fmt.Printf("defaultFloat64 (%T) = %f\n", defaultFloat64, defaultFloat64)
	fmt.Printf("defaultString (%T) = '%s'\n", defaultString, defaultString)
	fmt.Printf("defaultBool (%T) = %t\n", defaultBool, defaultBool)
}
