package main

import "fmt"

// Примеры использования функции func
func main() {
	round()
	add()
}

func round() {
	fmt.Println("Это функция round")
}

func add() {
	i := 2 + 5
	fmt.Println(i)
}