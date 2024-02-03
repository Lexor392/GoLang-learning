package main

import "fmt"

// Примеры использования ветвления if else
func main() {
	i := 2
	if i == 2 {
		fmt.Println("Переменная i ровна 2-ум")
	} else {
		fmt.Println("Переменная i больше либо меньше 2-ух")
	}

	s := 57
	if s >= 30 {
		fmt.Println("Переменная s больше 30-ти")
	} else if s <= 25 {
		fmt.Println("Переменная s меньше 25-ти")
	} else {
		fmt.Println("Переменная s находится в диапазоне от 26-ти до 29-ти")
	}
}