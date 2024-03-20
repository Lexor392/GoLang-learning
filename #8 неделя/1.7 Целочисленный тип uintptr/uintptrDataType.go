package main

import (
	"fmt"
	"unsafe"
)

/* Целочисленный тип uintptr */
func main() {
	var q uintptr = 0xc82000c2 // тип uintptr может хранить неизменяемые значения указателей или адресов в памяти.
	fmt.Println("Value of q:", q)

	var w = 456
	var e *int = &w // Преобразование в указатель

	// Преобразование указателя переменной e в uintptr
	uintptrAdd := uintptr(unsafe.Pointer(e))

	// Вывод адреса в формате 16-ричных чисел
	fmt.Printf("Adress of w: %x\n", uintptrAdd)

	// Преобразование обратно в указатель
	uintptrBack := (*int)(unsafe.Pointer(uintptrAdd))
	fmt.Printf("Value at adress: %d\n", *uintptrBack)
}

/* Тип uintptr это целочисленное представление адреса памяти.
Используется для выполнения косвенных арифметических операций
над небезопасными указателями для небезопасного доступа к памяти */
