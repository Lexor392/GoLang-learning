package main

import "fmt"

func main() {

	/* Игнорирование индекса в цикле по массиву */
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		fmt.Println(num)
	}

	fmt.Println("")

	/* Игнорирование значения элемента в цикле по слайсу */
	// используется для игнорирования значения элемента слайса
	numbers2 := [5]int{1, 2, 3, 4, 5}

	for i := range numbers2 {
		fmt.Println(i)
	}
}
