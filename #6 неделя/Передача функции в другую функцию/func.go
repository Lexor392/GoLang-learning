package main

import "fmt"

// Пример функции, принимающей другую функцию в качестве аргумента
func applyFunc(f func(int) int, x int) int {
	return f(x)
}

// Пример функции, которую мы будем передавать
func square(x int) int {
	return x * x
}

func main() {
	// Передаем функцию square в качестве аргумента функции applyFunction
	result := applyFunc(square, 5)
	fmt.Println(result) // Выведет: 25
}
