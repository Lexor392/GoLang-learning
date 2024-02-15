package main

import "fmt"

//Тип «nil» и когда его можно использовать
func main() {

	//Указатели
	var pointer *int
	if pointer == nil {
		fmt.Println("Указатель не указывает ни на один обьект")
	}

	//Функции
	var funcNil func() int
	if funcNil == nil {
		fmt.Println("Функция не инициализирована")
	}

	//Слайсы
	var sliceNil []int
	if sliceNil == nil {
		fmt.Println("Слайс не инициализирован")
	}

	//Другой пример (ошибка, где nil не будет использоваться)
	/*
		var number = 0

		if number == nil {
			fmt.Println("Переменная number имеет пустое значение")
		}
	*/
}
