package main

import "fmt"

func main() {
	fieldsStructure()
}

/* Пустая структура */
func emptyStructure() {

	type name struct{} // инициализация пустой структуры

	var emptyStruct name

	fmt.Println(emptyStruct)

}

/*Структура с полями*/
func fieldsStructure() {

	type Fields struct { //инициализация структуры с полями
		A int
		B int
		C int
	}

	fmt.Println(Fields{6, 4, 9}) // передача данных в структуру и вывод этой структуры в консоль

}