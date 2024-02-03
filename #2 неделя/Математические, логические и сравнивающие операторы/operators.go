package main

import "fmt"

func main() {
	// Математические операторы
	fmt.Println(1 + 1) //сложение

	fmt.Println(2 - 1) //вычитание

	fmt.Println(2 * 2) //умножение

	fmt.Println(4 / 2) //деление

	fmt.Println(14 % 5) //остаток от деления

	//Операторы сравнения
	fmt.Println(1 == 1) //равно

	fmt.Println(1 != 1) //не равно

	fmt.Println(2 > 1) //Больше

	fmt.Println(1 < 2) //Меньше

	fmt.Println(2 >= 1) //Больше или равно

	fmt.Println(1 <= 2) //Меньше или равно

	//Логические операторы
	fmt.Println((1 == 1) && (1 <= 2)) // && - оператор и

	fmt.Println((2 >= 1) || (2 < 3)) // || - оператор или

	fmt.Println(!(1 == 1)) // ! - оператор нет
}
