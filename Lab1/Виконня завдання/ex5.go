package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	// Коротка довідка про синоніми типів у Go.
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32 або int64, в залежності від платформи")
	fmt.Println("uint    - uint32 або uint64, в залежності від платформи")

	// Завдання: визначити розрядність платформи.
	// Розмір int/uint залежить від архітектури (32 або 64 біти).
	var testInt int
	var testUint uint
	
	fmt.Printf("\nРозрядність платформи:\n")
	// unsafe.Sizeof повертає розмір у байтах.
	fmt.Printf("int розмір: %d байт (%d біт)\n", unsafe.Sizeof(testInt), unsafe.Sizeof(testInt)*8)
	fmt.Printf("uint розмір: %d байт (%d біт)\n", unsafe.Sizeof(testUint), unsafe.Sizeof(testUint)*8)
	// runtime.GOARCH показує цільову архітектуру (наприклад, amd64).
	fmt.Printf("Архітектура: %s\n", runtime.GOARCH)
	
	// Якщо int займає 8 байт, система 64-бітна.
	if unsafe.Sizeof(testInt) == 8 {
		fmt.Println("Платформа: 64-бітна")
	} else {
		fmt.Println("Платформа: 32-бітна")
	}
}
