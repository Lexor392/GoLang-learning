package main

import "fmt"

// Примеры использования возвращаемых типов через return
func main() {
	x := 2
	y := 4

	fmt.Println(add(x, y))
	mult("This is result:", x, y)
}

func add(a, b int) int {
	return a + b
}

func mult(s string, a, b int) {
	result := a * b
	fmt.Println(s, result)
}
