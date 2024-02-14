package main

import "fmt"

func main() {
	sum := sameTypes(13, 5)
	fmt.Println(sum)
}

// возвращение одинаковых типов
func sameTypes(x, y int) int {
	s := (x + y) / 2
	return s + 2
}
