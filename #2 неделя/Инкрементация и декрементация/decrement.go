package main

import "fmt"

// Виды декрементации
func main() {
	x := 1
	x--       // постфиксный декремент
	x -= 1    // присваивание
	x = x - 1 // математическая операция
	fmt.Println(x)

	var y *int
	y = new(int)
	(*y)-- // декрементация переменной указателя (указатель не учил)
}
